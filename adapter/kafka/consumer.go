package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Consumer represents a Kafka consumer.
type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

// NewKafkaConsumer creates a new instance of Consumer.
func NewKafkaConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

// Consume reads a message from the Kafka topic and returns it.
func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return err
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		return err
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
