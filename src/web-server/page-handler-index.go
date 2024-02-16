package web_server

import (
	"log"
	"net/http"
	"troskove/db"
	"troskove/services"
)

type LayoutData struct {
	ExpenseTypes   []services.ExpenseType
	Expenses       []services.Expense
	ExpensesFilter db.GetExpensesQueryFilters
	TotalValue     float64
	AverageValue   float64
	LoggedIn       bool
}

func pageHandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := getIndexPageTemplate()

	if err != nil {
		handleError(w, err, "Error getting layout template", http.StatusInternalServerError)
		return
	}

	valid, err := services.GetAuthService().VerifyRequest(r)

	data := LayoutData{
		LoggedIn: valid,
	}

	if err != nil || !valid {
		log.Println("Unauthorized request")

		if err = tmpl.Execute(w, data); err != nil {
			handleError(w, err, "Error executing layout template", http.StatusInternalServerError)
			return
		}

		return
	}

	expenseTypesData, dbErr := services.GetExpenseTypeService().GetExpenseTypes()
	if dbErr != nil {
		handleError(w, dbErr, "Error getting expense types", http.StatusInternalServerError)
		return
	}

	filters := db.GetExpensesQueryFilters{
		TypeId:      r.URL.Query().Get("expenses-filter-expense-type-id"),
		StartDate:   r.URL.Query().Get("expenses-filter-start-date"),
		EndDate:     r.URL.Query().Get("expenses-filter-end-date"),
		BiggerThan:  r.URL.Query().Get("expenses-filter-bigger-than"),
		SmallerThan: r.URL.Query().Get("expenses-filter-smaller-than"),
	}

	expensesData, dbErr := services.GetExpensesService().GetExpenses(db.GetExpensesQueryOptions{
		Filters: filters,
	})

	if dbErr != nil {
		handleError(w, dbErr, "Error getting expenses", http.StatusInternalServerError)
		return
	}

	totalValue := services.GetExpensesService().GetTotalValue(expensesData)
	averageValue := services.GetExpensesService().GetAverageValue(expensesData)

	data = LayoutData{
		ExpenseTypes:   expenseTypesData,
		Expenses:       expensesData,
		ExpensesFilter: filters,
		TotalValue:     totalValue,
		AverageValue:   averageValue,
		LoggedIn:       valid,
	}

	if err = tmpl.Execute(w, data); err != nil {
		handleError(w, err, "Error executing layout template", http.StatusInternalServerError)
		return
	}
}
