package financial_management_application

import (
	financial_management_domain "troskove/finantial-management/domain"
)

type ITransactionService interface {
	Create(data CreateTransactionDTO) error
	GetAll(options financial_management_domain.FindAllOptions) ([]financial_management_domain.Transaction, error)
	UpdateTransaction(data UpdateTransactionDTO) error
	DeleteTransaction(id string) error
}

type TransactionService struct {
	transactionRepository         financial_management_domain.ITransactionRepository
	transactionCategoryRepository financial_management_domain.ITransactionCategoryRepository
}

func NewTransactionService(
	transactionRepository financial_management_domain.ITransactionRepository,
	transactionCategoryRepository financial_management_domain.ITransactionCategoryRepository,
) ITransactionService {
	return &TransactionService{transactionRepository, transactionCategoryRepository}
}
