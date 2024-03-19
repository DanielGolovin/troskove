package financial_management_infrastructure

import (
	"database/sql"
	financial_management_domain "troskove/finantial-management/domain"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(DB *sql.DB) financial_management_domain.ITransactionRepository {
	return &TransactionRepository{DB}
}
