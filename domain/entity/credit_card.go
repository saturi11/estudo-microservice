package entity

import (
	"errors"
	"regexp"
	"time"
)

// CreditCard represents a credit card entity.
type CreditCard struct {
	number          string // The credit card number.
	name            string // The name on the credit card.
	expirationMonth int    // The expiration month of the credit card.
	expirationYear  int    // The expiration year of the credit card.
	cvv             int    // The CVV (Card Verification Value) of the credit card.
}

// NewCreditCard creates a new CreditCard instance with the provided details.
// It validates the credit card details and returns an error if any validation fails.
func NewCreditCard(number, name string, expirationMonth, expirationYear, cvv int) (*CreditCard, error) {
	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

// IsValid checks if the credit card details are valid.
// It calls the individual validation methods and returns an error if any validation fails.
func (c *CreditCard) IsValid() error {
	err := c.validateNumber()
	if err != nil {
		return err
	}

	err = c.validateMonth()
	if err != nil {
		return err
	}

	err = c.validateYear()
	if err != nil {
		return err
	}
	err = c.validateCVV()
	if err != nil {
		return err
	}

	return nil
}

// validateNumber checks if the credit card number is valid.
// It uses a regular expression to match the pattern of a valid credit card number.
func (c *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

// validateMonth checks if the expiration month is valid.
// It ensures that the month is between 1 and 12 (inclusive).
func (c *CreditCard) validateMonth() error {
	if c.expirationMonth < 1 || c.expirationMonth > 12 {
		return errors.New("invalid expiration month")
	}
	return nil
}

// validateYear checks if the expiration year is valid.
// It ensures that the year is greater than or equal to the current year.
func (c *CreditCard) validateYear() error {
	if c.expirationYear >= time.Now().Year() {
		return nil
	}
	return errors.New("invalid expiration year")
}

// validateCVV checks if the CVV is valid.
// It ensures that the CVV is a 3-digit number.
func (c *CreditCard) validateCVV() error {
	if c.cvv < 100 || c.cvv > 999 {
		return errors.New("invalid CVV")
	}
	return nil
}
