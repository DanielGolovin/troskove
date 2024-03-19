package financial_management_infrastructure

func (r *TransactionCategoryRepository) Delete(id string) error {
	_, err := r.DB.Exec(`
			PRAGMA foreign_keys = ON;
			DELETE FROM transaction_category WHERE id = ?;
		`,
		id)

	return err
}
