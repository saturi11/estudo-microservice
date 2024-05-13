package entity

import "errors"

const (
	APPROVED = "approved"
	REJECTED = "rejected"
)

type Transactional struct {
	ID            string
	AccontId      string
	Amount        float64
	CreditCard    CreditCard
	Status        string
	ErrorMesssage string
}

func NewTransactional() *Transactional {
	return &Transactional{}
}

func (t *Transactional) Isvalid() error {
	if t.Amount > 1000 {
		return errors.New("insufficient funds in the account")
	}
	if t.Amount < 1 {
		return errors.New("transactional amount must be bigger than 1")
	}
	return nil
}

func (t *Transactional) SetCreditCard(creditCard *CreditCard) {
	t.CreditCard = *creditCard
}
