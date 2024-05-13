package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saturi11/gateway/adapter/broker/kafka"
	"github.com/saturi11/gateway/adapter/factory"
	"github.com/saturi11/gateway/adapter/presenter/transaction"
	"github.com/saturi11/gateway/usecase/process_transaction"
)

func main() {
	db, error := sql.Open("sqlite3", "test.db")
	if error != nil {
		log.Fatal(error)
	}
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	kafkaPresenter := transaction.NewTransactionalKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewKafkaConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionalDTOInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
