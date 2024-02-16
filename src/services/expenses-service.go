package services

import (
	"database/sql"
	"troskove/db"
)

var expensesServiceInstance IExpensesService

func GetExpensesService() IExpensesService {
	if expensesServiceInstance == nil {
		expensesServiceInstance = &ExpensesService{
			DB: db.GetDBConnection(),
		}
	}

	return expensesServiceInstance
}

type Expense struct {
	ID       string
	Value    float64
	Date     string
	TypeID   string
	TypeName string
}

type ExpensesService struct {
	DB *sql.DB
	IExpensesService
}

type IExpensesService interface {
	CreateExpense(expense db.InsertExpenseDTO) error
	GetExpenses(options db.GetExpensesQueryOptions) ([]Expense, error)
	UpdateExpense(id string, expense db.UpdateExpenseDTO) error
	DeleteExpense(id string) error
	GetTotalValue([]Expense) float64
	GetAverageValue([]Expense) float64
}

func (es *ExpensesService) CreateExpense(expense db.InsertExpenseDTO) error {
	_, err := db.InsertExpense(es.DB, expense)
	return err
}

func (es *ExpensesService) GetExpenses(options db.GetExpensesQueryOptions) ([]Expense, error) {

	if options.Limit == 0 || options.Limit > 250 {
		options.Limit = 250
	}

	rows, err := db.GetExpensesQuery(es.DB, options)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	expenses, err := parseExpenses(rows)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (es *ExpensesService) UpdateExpense(id string, expense db.UpdateExpenseDTO) error {
	_, err := db.UpdateExpense(es.DB, id, expense)
	return err
}

func (es *ExpensesService) DeleteExpense(id string) error {
	_, err := db.DeleteExpense(es.DB, id)
	return err
}

func (es *ExpensesService) GetTotalValue(expenses []Expense) float64 {
	var total float64

	for _, expense := range expenses {
		total += expense.Value
	}

	return total
}

func (es *ExpensesService) GetAverageValue(expenses []Expense) float64 {
	total := es.GetTotalValue(expenses)
	return total / float64(len(expenses))
}

func parseExpenses(rows *sql.Rows) ([]Expense, error) {
	var expenses []Expense

	for rows.Next() {
		var expense Expense

		typeName := sql.NullString{}

		err := rows.Scan(&expense.ID, &expense.Value, &expense.Date, &expense.TypeID, &typeName)

		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
