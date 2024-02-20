package db

import "database/sql"

func DeleteExpense(db *sql.DB, id string) (sql.Result, error) {
	query := `DELETE FROM expenses WHERE id = $1;`

	return db.Exec(query, id)
}
