package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCreditCardNumber tests the NewCreditCard function with an invalid credit card number.
func TestCreditCardNumber(t *testing.T) {

	_, err := NewCreditCard("1000000000000000", "Jose da silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da silva", 12, 2024, 123)
	assert.Nil(t, err)
}

// TestCreditCardExpirationMonth tests the NewCreditCard function with an invalid expiration month.
func TestCreditCardExpirationMonth(t *testing.T) {

	_, err := NewCreditCard("4193523830170205", "Jose da silva", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da silva", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da silva", 12, 2024, 123)
	assert.Nil(t, err)
}

// TestCreditCardExpirationYear tests the NewCreditCard function with an invalid expiration year.
func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("4193523830170205", "Jose da silva", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da silva", 12, 2025, 123)
	assert.Nil(t, err)
}

// TestCreditCardCvv tests the NewCreditCard function with an invalid cvv.
func TestCreditCardCvv(t *testing.T) {

	_, err := NewCreditCard("4193523830170205", "Jose da silva", 12, 2025, 123)
	assert.Nil(t, err)

	_, err = NewCreditCard("4193523830170205", "Jose da silva", 12, 2025, 1234)
	assert.Equal(t, "invalid cvv", err.Error())
}
