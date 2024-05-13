package factory

import (
	"database/sql"

	repo "github.com/saturi11/gateway/adapter/repository"
	"github.com/saturi11/gateway/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionalRepository {
	return repo.NewTransactionalRepositoryDb(r.DB)
}
