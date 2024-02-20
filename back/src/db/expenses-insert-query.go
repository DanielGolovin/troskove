package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type InsertExpenseDTO struct {
	Value  float64
	TypeID string
	Date   string
}

func (expense InsertExpenseDTO) Validate() error {
	if expense.Value == 0 {
		return errors.New("Value is required")
	}
	if expense.TypeID == "" {
		return errors.New("Type ID is required")
	}
	return nil
}

func InsertExpense(db *sql.DB, expense InsertExpenseDTO) (sql.Result, error) {
	sqlStmt := `
			INSERT INTO expenses (id, value, date, type_id)
			VALUES (?, ?, ?, ?);
			`

	id := uuid.New().String()
	date := expense.Date

	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}

	return db.Exec(sqlStmt, id, expense.Value, date, expense.TypeID)
}
