package process_transaction
import (
"testing"
"time"
"github.com/saturi11/getway/domain/repository"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	transactionalDTOInput := TransactionalDTOInput{
		ID: "1",
		AccontId: "1",
		TestCreditCardNumber: "1000000000000000",
		CreditCardName: "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		CreditCardCvv: 123,
		Amount: 500,
	}
	expectedOutput := TransactionalDTOOutput{
		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := repository.NewMockransactionRepository(ctrl)
	repositoryMock.EXPECT().
	Insert(inpiut.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage)
	.return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}