package repository

import (
	"testing"

	"github.com/saturi11/gateway/adapter/repository/fixture"
	"github.com/saturi11/gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionalDbInsert(t *testing.T) {
	migrationsDir := "fixture/sql"
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionalRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
