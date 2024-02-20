package db

import (
	"database/sql"
	"errors"
)

type UpdateExpenseTypeDTO struct {
	Name string
}

func (expenseType UpdateExpenseTypeDTO) Validate() error {
	if expenseType.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}

func UpdateExpenseType(db *sql.DB, id string, expenseType UpdateExpenseTypeDTO) (sql.Result, error) {
	query := `UPDATE expense_types SET name = $1 WHERE id = $2;`

	return db.Exec(query, expenseType.Name, id)
}
