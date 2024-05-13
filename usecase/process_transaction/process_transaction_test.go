package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_broker "github.com/saturi11/gateway/adapter/broker/mock"
	"github.com/saturi11/gateway/domain/entity"
	mock_repository "github.com/saturi11/gateway/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

// TestProcessTransaction_ExecuteInvalidCreditCard tests the execution of the ProcessTransaction use case
// when an invalid credit card is provided.
func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionalDTOInput{
		ID:                        "1",
		AccontId:                  "1",
		TestCreditCardNumber:      "1000000000000000",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCvv:             123,
		Amount:                    500,
	}

	expectedOutput := TransactionalDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_repository.NewMockRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccontId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducer(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_result")
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

// TestProcessTransaction_ExecuteRejectedTransaction tests the execution of the ProcessTransaction use case
// when a transaction is rejected due to insufficient funds in the account.
func TestProcessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	input := TransactionalDTOInput{
		ID:                        "1",
		AccontId:                  "1",
		TestCreditCardNumber:      "4193523830170205",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCvv:             123,
		Amount:                    1200,
	}

	expectedOutput := TransactionalDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "insufficient funds in the account",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_repository.NewMockRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccontId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducer(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_result")
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

// TestProcessTransaction_ExecuteApprovedTransaction tests the execution of the ProcessTransaction use case
// when a transaction is approved.
func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionalDTOInput{
		ID:                        "1",
		AccontId:                  "1",
		TestCreditCardNumber:      "4193523830170205",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCvv:             123,
		Amount:                    900,
	}

	expectedOutput := TransactionalDTOOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_repository.NewMockRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccontId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducer(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_result")
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
