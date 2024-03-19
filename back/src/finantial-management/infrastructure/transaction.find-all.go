package financial_management_infrastructure

import (
	"fmt"
	"strings"
	financial_management_domain "troskove/finantial-management/domain"
)

var allowedColumns = map[string]string{
	"date":  "transactions.date",
	"value": "transactions.amount",
	"type":  "transaction_category.name",
	"id":    "transactions.id",
}

var allowedDirections = map[string]string{
	"asc":  "ASC",
	"desc": "DESC",
}

func validateOrderBy(orderBy string) string {
	safeOrderBy, ok := allowedColumns[orderBy]
	if !ok {
		return "transactions.date"
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

func getExpsensesQueryFilters(filters financial_management_domain.FindAllFilters) (string, []interface{}) {
	var args []interface{}

	stmtlines := []string{}
	q := "WHERE "

	if filters.CategoryID != "" {
		stmtlines = append(stmtlines, "transactions.category_id = ?")
		args = append(args, filters.CategoryID)
	}

	if filters.StartDate != "" {
		stmtlines = append(stmtlines, "transactions.date >= ?")
		args = append(args, filters.StartDate)
	}

	if filters.EndDate != "" {
		stmtlines = append(stmtlines, "transactions.date <= ?")
		args = append(args, filters.EndDate)
	}

	if filters.BiggerThan != "" {
		stmtlines = append(stmtlines, "transactions.amount >= ?")
		args = append(args, filters.BiggerThan)
	}

	if filters.SmallerThan != "" {
		stmtlines = append(stmtlines, "transactions.amount <= ?")
		args = append(args, filters.SmallerThan)
	}

	q += strings.Join(stmtlines, " AND ")

	if q == "WHERE " {
		return "", args
	}

	return q, args
}

func (repo *TransactionRepository) FindAll(options financial_management_domain.FindAllOptions) ([]financial_management_domain.Transaction, error) {
	safeOrderBy := validateOrderBy(options.OrderBy)
	safeOrderDir := validateOrderDirection(options.OrderDir)
	filtersStmt, filtersArgs := getExpsensesQueryFilters(options.Filters)

	sqlStmt := fmt.Sprintf(`
	    SELECT transactions.id, transactions.amount, transactions.date, transactions.category_id, transaction_category.name FROM transactions
	    JOIN transaction_category ON transactions.category_id = transaction_category.id
		%s
	    ORDER BY %s %s
	    LIMIT ? OFFSET ?
	`, filtersStmt, safeOrderBy, safeOrderDir)

	args := append(filtersArgs, options.Limit, options.Offset)

	rows, err := repo.DB.Query(sqlStmt, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []financial_management_domain.Transaction

	for rows.Next() {
		var transaction financial_management_domain.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.Date,
			&transaction.Category.ID,
			&transaction.Category.Name,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
