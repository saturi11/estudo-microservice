package entity
import (
    "testing"
    "github.com/stretchr/testify/assert"
	
)

func TestTransactional_IsValid(t *testing.T) {
	transactional := NewTransactional()
		transactional.ID = "1"
		transactional.Amount = 500
		transactional.AccontId = "1"
		assert.Nil(t,transactional.Isvalid())
	
}

func TestTransactional_IsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transactional := NewTransactional()
	transactional.ID = "1"
	transactional.Amount = 1001
	transactional.AccontId = "1"
	err := transactional.Isvalid()
	assert.Error(t,err)
	assert.Equal(t,"transactional amount must be less than 1000",err.Error())
}


func TestTransactional_IsNotValidWithAmountLassThen1(t *testing.T) {
	transactional := NewTransactional()
	transactional.ID = "1"
	transactional.Amount = 1001
	transactional.AccontId = "1"
	err := transactional.Isvalid()
	assert.Error(t,err)
	assert.Equal(t,"transactional amount must be bigger than 1",err.Error())
}