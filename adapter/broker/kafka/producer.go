package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saturi11/gateway/adapter/presenter"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{
		ConfigMap: configMap,
		Presenter: presenter,
	}
}

// Publish sends the given message to the specified topic in Kafka.
// The `msg` parameter represents the message to be published.
// The `key` parameter is an optional key used for message partitioning.
// The `topic` parameter specifies the topic to which the message should be published.
// Returns an error if the message fails to be published.
// Example usage: err := producer.Publish("Hello, Kafka!", []byte("key"), "my-topic")
// Note: This function assumes that the Kafka producer is already initialized and connected.
func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}
	err = p.Presenter.Bind(msg)
	if err != nil {
		return err
	}
	presenterMsg, err := p.Presenter.Show()
	if err != nil {
		return err
	}
	message := ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: -1},
		Value:          presenterMsg,
		Key:            key,
	}

	err = producer.Produce(&message, nil)
	if err != nil {
		return err
	}
	return nil
}
