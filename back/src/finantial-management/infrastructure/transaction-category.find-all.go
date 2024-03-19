package financial_management_infrastructure

import (
	financial_management_domain "troskove/finantial-management/domain"
)

func (r *TransactionCategoryRepository) FindAll() ([]financial_management_domain.TransactionCategory, error) {
	rows, err := r.DB.Query("SELECT id, name FROM transaction_category")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []financial_management_domain.TransactionCategory

	for rows.Next() {
		var category financial_management_domain.TransactionCategory
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
