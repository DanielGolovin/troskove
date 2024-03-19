package financial_management_infrastructure

import (
	"database/sql"
	financial_management_domain "troskove/finantial-management/domain"
)

type TransactionCategoryRepository struct {
	DB *sql.DB
}

func NewTransactionCategoryRepository(db *sql.DB) financial_management_domain.ITransactionCategoryRepository {
	return &TransactionCategoryRepository{db}
}
