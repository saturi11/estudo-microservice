package transaction

import (
	"encoding/json"

	"github.com/saturi11/gateway/usecase/process_transaction"
)

// KafkaPresenter is a presenter that formats transactional data for Kafka.
type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

// NewTransactionalKafkaPresenter creates a new instance of KafkaPresenter.
func NewTransactionalKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

// Bind binds the input data to the KafkaPresenter.
func (t *KafkaPresenter) Bind(input interface{}) error {
	t.ID = input.(process_transaction.TransactionalDTOOutput).ID
	t.Status = input.(process_transaction.TransactionalDTOOutput).Status
	t.ErrorMessage = input.(process_transaction.TransactionalDTOOutput).ErrorMessage
	return nil
}

// Show returns the JSON representation of the KafkaPresenter.
func (t *KafkaPresenter) Show() ([]byte, error) {
	J, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return J, nil
}
