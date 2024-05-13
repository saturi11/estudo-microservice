package transaction

import (
	"encoding/json"

	"github.com/saturi11/gateway/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionalKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (t *KafkaPresenter) Bind(input interface{}) error {
	t.ID = input.(process_transaction.TransactionalDTOOutput).ID
	t.Status = input.(process_transaction.TransactionalDTOOutput).Status
	t.ErrorMessage = input.(process_transaction.TransactionalDTOOutput).ErrorMessage
	return nil
}

func (t *KafkaPresenter) Show() ([]byte, error) {
	J, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return J, nil
}
