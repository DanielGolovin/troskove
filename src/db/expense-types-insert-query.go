package db

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type InsertExpenseTypeDTO struct {
	Name string
}

func (expenseType InsertExpenseTypeDTO) Validate() error {
	if expenseType.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}

func InsertExpenseType(db *sql.DB, data InsertExpenseTypeDTO) (sql.Result, error) {
	sqlStmt := `
			INSERT INTO expense_types (id, name)
			VALUES (?, ?);
			`
	id := uuid.New().String()

	return db.Exec(sqlStmt, id, data.Name)
}
