package entity

import "errors"

// Constants for transaction status
const (
	APPROVED = "approved"
	REJECTED = "rejected"
)

// Transactional represents a transaction entity
type Transactional struct {
	ID            string     // Unique identifier for the transaction
	AccontId      string     // Identifier for the associated account
	Amount        float64    // Transaction amount
	CreditCard    CreditCard // Credit card information associated with the transaction
	Status        string     // Current status of the transaction (approved or rejected)
	ErrorMesssage string     // Error message in case of a rejected transaction
}

// NewTransactional creates a new instance of Transactional
func NewTransactional() *Transactional {
	return &Transactional{}
}

// Isvalid checks if the transaction is valid
func (t *Transactional) Isvalid() error {
	if t.Amount > 1000 {
		return errors.New("insufficient funds in the account")
	}
	if t.Amount < 1 {
		return errors.New("transactional amount must be bigger than 1")
	}
	return nil
}

// SetCreditCard sets the credit card information for the transaction
func (t *Transactional) SetCreditCard(creditCard *CreditCard) {
	t.CreditCard = *creditCard
}
