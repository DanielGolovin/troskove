package financial_management_infrastructure

import (
	financial_management_domain "troskove/finantial-management/domain"
)

func (r *TransactionCategoryRepository) FindById(id string) (*financial_management_domain.TransactionCategory, error) {
	var transactionCategory financial_management_domain.TransactionCategory

	err := r.DB.QueryRow(
		"SELECT id, name FROM transaction_category WHERE id = $1",
		id,
	).Scan(
		&transactionCategory.ID,
		&transactionCategory.Name,
	)

	if err != nil {
		return nil, err
	}

	return &transactionCategory, nil
}
