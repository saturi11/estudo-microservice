package repository

type TransactionalRepository interface {
	Insert(id string, account string, amount float64, status string, ErrorMesssage string) error
}
