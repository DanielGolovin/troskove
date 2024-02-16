package db

import "database/sql"

func GetExpenseTypesQuery(db *sql.DB) (*sql.Rows, error) {
	sqlStmt := `
		SELECT id, name FROM expense_types;
	`

	return db.Query(sqlStmt)
}
