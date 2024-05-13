package process_transaction

import (
	"github.com/saturi11/gateway/domain/entity"
	"github.com/saturi11/gateway/domain/repository"
)

// Processtransaction is a struct that represents the process transaction use case.
type Processtransaction struct {
	Repository repository.TransactionalRepository
}

// NewProcessTransaction is a function that creates a new instance of the Processtransaction struct.
func NewProcessTransaction(repository repository.TransactionalRepository) *Processtransaction {
	return &Processtransaction{Repository: repository}
}

// Execute is a method that executes the process transaction use case.
// It takes an input of type TransactionalDTOInput and returns a TransactionalDTOOutput and an error.
func (p *Processtransaction) Execute(input TransactionalDTOInput) (TransactionalDTOOutput, error) {
	transactional := entity.NewTransactional()
	transactional.ID = input.ID
	transactional.AccontId = input.AccontId
	transactional.Amount = input.Amount

	cc, invalicCC := entity.NewCreditCard(input.TestCreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCvv)
	if invalicCC != nil {
		return p.rejectdTransaction(transactional, invalicCC)
	}

	transactional.SetCreditCard(cc)
	invalidTransactional := transactional.Isvalid()
	if invalidTransactional != nil {
		return p.rejectdTransaction(transactional, invalidTransactional)
	}

	return p.approvedTransaction(input, transactional)
}

// approvedTransaction is a method that handles the approved transaction case.
// It takes an input of type TransactionalDTOInput and a pointer to an entity.Transactional,
// and returns a TransactionalDTOOutput and an error.
func (p *Processtransaction) approvedTransaction(input TransactionalDTOInput, transactional *entity.Transactional) (TransactionalDTOOutput, error) {
	err := p.Repository.Insert(transactional.ID, transactional.AccontId, transactional.Amount, entity.APPROVED, "")
	if err != nil {
		return TransactionalDTOOutput{}, err
	}
	output := TransactionalDTOOutput{
		ID:           transactional.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}
	return output, nil
}

// rejectdTransaction is a method that handles the rejected transaction case.
// It takes a pointer to an entity.Transactional and an error,
// and returns a TransactionalDTOOutput and an error.
func (p *Processtransaction) rejectdTransaction(transactional *entity.Transactional, invalidTransactional error) (TransactionalDTOOutput, error) {
	err := p.Repository.Insert(transactional.ID, transactional.AccontId, transactional.Amount, entity.REJECTED, invalidTransactional.Error())
	if err != nil {
		return TransactionalDTOOutput{}, err
	}
	output := TransactionalDTOOutput{
		ID:           transactional.ID,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTransactional.Error(),
	}
	return output, nil
}
