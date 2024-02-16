package db

import "database/sql"

func DeleteExpenseType(db *sql.DB, id string) (sql.Result, error) {
	query := `
	PRAGMA foreign_keys = ON;
	DELETE FROM expense_types WHERE id = $1;`

	return db.Exec(query, id)
}
