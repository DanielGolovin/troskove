package financial_management_application

import (
	financial_management_domain "troskove/finantial-management/domain"
)

type ITransactionCategoryService interface {
	Create(name string) error
	GetAll() ([]financial_management_domain.TransactionCategory, error)
	Update(id string, dto UpdateTransactionCategoryDTO) error
	Delete(id string) error
}

type UpdateTransactionCategoryDTO struct {
	Name string
}

type TransactionCategoryService struct {
	transactionCategoryRepository financial_management_domain.ITransactionCategoryRepository
}

func NewTransactionCategoryService(transactionCategoryRepository financial_management_domain.ITransactionCategoryRepository) ITransactionCategoryService {
	return &TransactionCategoryService{transactionCategoryRepository}
}

func (service *TransactionCategoryService) Create(name string) error {
	transactionCategory, err := financial_management_domain.NewTransactionCategory(name)

	if err != nil {
		return err
	}

	return service.transactionCategoryRepository.Create(*transactionCategory)
}

func (service *TransactionCategoryService) GetAll() ([]financial_management_domain.TransactionCategory, error) {
	return service.transactionCategoryRepository.FindAll()
}

func (service *TransactionCategoryService) Update(id string, data UpdateTransactionCategoryDTO) error {
	transactionCategory, err := financial_management_domain.NewTransactionCategory(data.Name)

	if err != nil {
		return err
	}

	return service.transactionCategoryRepository.Update(id, *transactionCategory)
}

func (service *TransactionCategoryService) Delete(id string) error {
	return service.transactionCategoryRepository.Delete(id)
}
