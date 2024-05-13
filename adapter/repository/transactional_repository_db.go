package repository

import (
	"database/sql"
	"time"
)

type TransactionalRepositoryDb struct {
	db *sql.DB
}

func NewTransactionalRepositoryDb(db *sql.DB) *TransactionalRepositoryDb {
	return &TransactionalRepositoryDb{db: db}
}

func (r *TransactionalRepositoryDb) Insert(id string, account string, amount float64, status string, ErrorMesssage string) error {
	stmt, err := r.db.Prepare(`INSERT INTO transactions (id, account, amount, status, error_message, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		id,
		account,
		amount,
		status,
		ErrorMesssage,
		time.Now(),
		time.Now())
	if err != nil {
		return err
	}
	return nil
}
