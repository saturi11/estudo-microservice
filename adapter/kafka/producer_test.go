// Package kafka provides functionality for interacting with Kafka.
package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saturi11/gateway/adapter/presenter/transaction"
	"github.com/saturi11/gateway/domain/entity"
	"github.com/saturi11/gateway/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

// TestProducerPublish is a unit test for the Publish method of the KafkaProducer.
func TestProducerPublish(t *testing.T) {
	// Define the expected output of the Publish method.
	expectedOutput := process_transaction.TransactionalDTOOutput{
		ID:           "123",
		Status:       entity.REJECTED,
		ErrorMessage: "insufficient funds in the account",
	}

	// Create a Kafka configuration map.
	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	// Create a new KafkaProducer instance.
	produce := NewKafkaProducer(&configMap, transaction.NewTransactionalKafkaPresenter())

	// Publish the expected output to the Kafka topic.
	err := produce.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
