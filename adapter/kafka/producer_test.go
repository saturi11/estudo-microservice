package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saturi11/gateway/adapter/presenter/transaction"
	"github.com/saturi11/gateway/domain/entity"
	"github.com/saturi11/gateway/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionalDTOOutput{
		ID:           "123",
		Status:       entity.REJECTED,
		ErrorMessage: "insufficient funds in the account",
	}

	//outputJson,_ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	produce := NewKafkaProducer(&configMap, transaction.NewTransactionalKafkaPresenter())
	err := produce.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)

}
