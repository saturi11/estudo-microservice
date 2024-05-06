package process_transaction
import "estudo-microservice/domain/repository"

type Processtransaction struct {	
    Repository repository.TransactionalRepository 

func NewProcessTransaction(repository repository.TransactionalRepository) *Process_transaction {
    return &Processtransaction{Repository: repository}
}



func (p *ProcessTransaction) Execute(input TransactionalDTOInput) (TransactionalDTOOutput, error) {
	transactional := entity.NewTransactional()
	transactional.ID = input.ID
	transactional.AccontId = input.AccontId
	transactional.Amount = input.Amount
	_, invalicCC := entity.NewCreditCard(input.TestCreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCvv)
	if invalicCC != nil {
		err := p.Repository.Insert(transactional.ID, transactional.AccontId, transactional.Amount, entity.REJECTED, invalicCC.Error())
		if err != nil {
			return TransactionalDTOOutput{}, err
		}
	}
	return TransactionalDTOOutput{}, nil
}