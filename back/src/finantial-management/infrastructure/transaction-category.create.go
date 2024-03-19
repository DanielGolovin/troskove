package financial_management_infrastructure

import (
	financial_management_domain "troskove/finantial-management/domain"
)

func (r *TransactionCategoryRepository) Create(category financial_management_domain.TransactionCategory) error {
	_, err := r.DB.Exec("INSERT INTO transaction_category (id, name) VALUES (?, ?)", category.ID, category.Name)

	return err
}
