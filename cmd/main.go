package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	_ "github.com/mattn/go-sqlite3"
	"github.com/saturi11/gateway/adapter/broker/kafka"
	"github.com/saturi11/gateway/adapter/factory"
	"github.com/saturi11/gateway/adapter/presenter/transaction"
	"github.com/saturi11/gateway/usecase/process_transaction"
)

func main() {
	// Open a connection to the SQLite database
	db, error := sql.Open("sqlite3", "test.db")
	if error != nil {
		log.Fatal(error)
	}

	// Create a repository factory using the database connection
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// Configure the Kafka producer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	kafkaPresenter := transaction.NewTransactionalKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	// Create a channel to receive Kafka messages
	var msgChan = make(chan *ckafka.Message)

	// Configure the Kafka consumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewKafkaConsumer(configMapConsumer, topics)

	// Start consuming messages from Kafka in a separate goroutine
	go consumer.Consume(msgChan)

	// Create the use case for processing transactions
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	// Process incoming messages from Kafka
	for msg := range msgChan {
		var input process_transaction.TransactionalDTOInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
