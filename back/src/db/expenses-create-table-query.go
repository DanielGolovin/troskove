package db

import "database/sql"

type Expense struct {
	ID     string
	Value  float64
	Date   string
	TypeID string
}

func createExpensesTableQuery(db *sql.DB) (sql.Result, error) {
	sqlStmt := `
			CREATE TABLE IF NOT EXISTS expenses (
					id TEXT NOT NULL UNIQUE PRIMARY KEY,
					value REAL NOT NULL,
					date TEXT NOT NULL,
					type_id TEXT NOT NULL,
					FOREIGN KEY(type_id) REFERENCES expense_types(id) ON DELETE RESTRICT ON UPDATE CASCADE
			);
			
			CREATE INDEX idx_expenses_date ON expenses(date DESC);
			`

	return db.Exec(sqlStmt)
}
