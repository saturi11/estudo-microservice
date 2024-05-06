package repository

type Repository interface {
	Insert(id string, account string, amount float64, status string, ErrorMesssage string ) error
}