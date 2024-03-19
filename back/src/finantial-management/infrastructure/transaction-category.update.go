package financial_management_infrastructure

import financial_management_domain "troskove/finantial-management/domain"

func (r *TransactionCategoryRepository) Update(id string, category financial_management_domain.TransactionCategory) error {
	_, err := r.DB.Exec("UPDATE transaction_category SET name = ? WHERE id = ?", category.Name, id)

	return err
}
