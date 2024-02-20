package db

import (
	"database/sql"
	"fmt"
	"strings"
)

var allowedColumns = map[string]string{
	"date":  "expenses.date",
	"value": "expenses.value",
	"type":  "expense_types.name",
	"id":    "expenses.id",
}

var allowedDirections = map[string]string{
	"asc":  "ASC",
	"desc": "DESC",
}

func validateOrderBy(orderBy string) string {
	safeOrderBy, ok := allowedColumns[orderBy]
	if !ok {
		return "expenses.date"
	}
	return safeOrderBy
}

func validateOrderDirection(orderDir string) string {
	safeOrderDir, ok := allowedDirections[strings.ToLower(orderDir)]
	if !ok {
		return "DESC"
	}
	return safeOrderDir
}

type GetExpensesQueryOptions struct {
	OrderBy  string
	OrderDir string
	Limit    int
	Offset   int
	Filters  GetExpensesQueryFilters
}

type GetExpensesQueryFilters struct {
	TypeId      string
	StartDate   string
	EndDate     string
	BiggerThan  string
	SmallerThan string
}

func getExpsensesQueryFilters(filters GetExpensesQueryFilters) (string, []interface{}) {
	var args []interface{}

	stmtlines := []string{}
	q := "WHERE "

	if filters.TypeId != "" {
		stmtlines = append(stmtlines, "expenses.type_id = ?")
		args = append(args, filters.TypeId)
	}

	if filters.StartDate != "" {
		stmtlines = append(stmtlines, "expenses.date >= ?")
		args = append(args, filters.StartDate)
	}

	if filters.EndDate != "" {
		stmtlines = append(stmtlines, "expenses.date <= ?")
		args = append(args, filters.EndDate)
	}

	if filters.BiggerThan != "" {
		stmtlines = append(stmtlines, "expenses.value >= ?")
		args = append(args, filters.BiggerThan)
	}

	if filters.SmallerThan != "" {
		stmtlines = append(stmtlines, "expenses.value <= ?")
		args = append(args, filters.SmallerThan)
	}

	q += strings.Join(stmtlines, " AND ")

	if q == "WHERE " {
		return "", args
	}

	return q, args
}

func GetExpensesQuery(db *sql.DB, options GetExpensesQueryOptions) (*sql.Rows, error) {
	safeOrderBy := validateOrderBy(options.OrderBy)
	safeOrderDir := validateOrderDirection(options.OrderDir)
	filtersStmt, filtersArgs := getExpsensesQueryFilters(options.Filters)

	sqlStmt := fmt.Sprintf(`
	    SELECT expenses.id, expenses.value, expenses.date, expenses.type_id, expense_types.name FROM expenses
	    JOIN expense_types ON expenses.type_id = expense_types.id
		%s
	    ORDER BY %s %s
	    LIMIT ? OFFSET ?
	`, filtersStmt, safeOrderBy, safeOrderDir)

	args := append(filtersArgs, options.Limit, options.Offset)

	rows, err := db.Query(sqlStmt, args...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
