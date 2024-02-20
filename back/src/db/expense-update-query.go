package db

import (
	"database/sql"
	"errors"
)

type UpdateExpenseDTO struct {
	Value  float64
	Date   string
	TypeID string
}

func (expense UpdateExpenseDTO) Validate() error {
	if expense.Value == 0 {
		return errors.New("Value is required")
	}
	if expense.TypeID == "" {
		return errors.New("Type ID is required")
	}
	if expense.Date == "" {
		return errors.New("Date is required")
	}
	return nil
}

func UpdateExpense(db *sql.DB, id string, expense UpdateExpenseDTO) (sql.Result, error) {
	query := `UPDATE expenses SET value = $1, date = $2, type_id = $3 WHERE id = $4`

	return db.Exec(query, expense.Value, expense.Date, expense.TypeID, id)
}
