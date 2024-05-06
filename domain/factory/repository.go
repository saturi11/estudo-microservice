package factory

import "github.com/saturi11/gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionalRepository() repository.TransactionalRepository
}