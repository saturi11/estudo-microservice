package entity

import (
    "testing"
    "github.com/stretchr/testify/assert"
	"time"
)

func TestCreditCardNumber(t *testing.T) {

  _, err := NewCreditCard("1000000000000000","Jose da silva",12,2024,123)
  assert.Equal(t, "invalid credit card number", err.Error())


  _, err = NewCreditCard("4193523830170205","Jose da silva",12,2024,123)
  assert.Nil(t,err)
}

func TestCreditCardExpirationMonth(t *testing.T) {

  _, err := NewCreditCard("4193523830170205","Jose da silva",0,2024,123)
  assert.Equal(t, "invalid expiration month", err.Error())

  _, err = NewCreditCard("4193523830170205","Jose da silva",13,2024,123)
  assert.Equal(t, "invalid expiration month", err.Error())

  _, err = NewCreditCard("4193523830170205","Jose da silva",12,2024,123)
  assert.Nil(t,err)
}


func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1,0,0)
  _, err := NewCreditCard("4193523830170205","Jose da silva",12,lastYear.Year(),123)
  assert.Equal(t, "invalid expiration year", err.Error())

  _, err = NewCreditCard("4193523830170205","Jose da silva",12,2025,123)
  assert.Nil(t,err)
}

func TestCreditCardCvv(t *testing.T) {

  _, err := NewCreditCard("4193523830170205","Jose da silva",12,2025,123)
  assert.Nil(t,err)

  _, err = NewCreditCard("4193523830170205","Jose da silva",12,2025,1234)
  assert.Equal(t, "invalid cvv", err.Error())
}