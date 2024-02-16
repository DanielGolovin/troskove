package db

import "database/sql"

type ExpenseType struct {
	ID   string
	Name string
}

func createExpenseTypesTableQuery(db *sql.DB) (sql.Result, error) {
	sqlStmt := `
			CREATE TABLE IF NOT EXISTS expense_types (
					id TEXT NOT NULL UNIQUE PRIMARY KEY,
					name TEXT NOT NULL UNIQUE
			);
			`

	return db.Exec(sqlStmt)
}
