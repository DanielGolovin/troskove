package financial_management_infrastructure

import "database/sql"

func InitTransactionTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id TEXT NOT NULL UNIQUE PRIMARY KEY,
			amount REAL NOT NULL,
			date TEXT NOT NULL,
			category_id TEXT NOT NULL,
			FOREIGN KEY(category_id) REFERENCES transaction_category(id) ON DELETE RESTRICT ON UPDATE CASCADE
		);
	`)

	if err != nil {
		return err
	}

	return nil
}

func InitTransactionCategoryTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS transaction_category (
			id TEXT NOT NULL UNIQUE PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
