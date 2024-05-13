// This file contains the unit test for the TransactionalRepositoryDb type in the repository package.
// It tests the functionality of the Insert method.

package repository

import (
	"testing"

	"github.com/saturi11/gateway/adapter/repository/fixture"
	"github.com/saturi11/gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

// TestTransactionalDbInsert is a unit test function that tests the Insert method of the TransactionalRepositoryDb type.
// It sets up the necessary fixtures, inserts a transaction into the database, and asserts that no error occurred.
func TestTransactionalDbInsert(t *testing.T) {
	migrationsDir := "fixture/sql"
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionalRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
