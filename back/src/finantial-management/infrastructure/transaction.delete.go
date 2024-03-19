package financial_management_infrastructure

func (repository *TransactionRepository) Delete(id string) error {
	_, err := repository.DB.Exec(`
			DELETE FROM transactions WHERE id = $1;
		`,
		id,
	)

	return err
}
