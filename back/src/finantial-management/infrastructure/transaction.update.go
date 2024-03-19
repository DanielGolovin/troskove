package financial_management_infrastructure

import (
	financial_management_domain "troskove/finantial-management/domain"
)

func (repository *TransactionRepository) Update(id string, transaction financial_management_domain.Transaction) error {
	_, err := repository.DB.Exec(`
		UPDATE transactions 
		SET amount = $1, date = $2, category_id = $3 WHERE id = $4
	`,
		transaction.Amount, transaction.Date, transaction.Category.ID, id,
	)
	return err
}
