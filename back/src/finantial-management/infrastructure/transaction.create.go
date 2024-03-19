package financial_management_infrastructure

import (
	financial_management_domain "troskove/finantial-management/domain"
)

func (repository *TransactionRepository) Create(transaction financial_management_domain.Transaction) error {
	_, err := repository.DB.Exec(
		"INSERT INTO transactions (id, amount, date, category_id) VALUES ($1, $2, $3, $4)",
		transaction.ID, transaction.Amount, transaction.Date, transaction.Category.ID,
	)

	return err
}
