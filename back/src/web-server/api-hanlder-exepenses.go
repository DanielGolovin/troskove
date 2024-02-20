package web_server

import (
	"encoding/json"
	"net/http"
	"troskove/db"
	"troskove/services"
)

func expensesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetExpenses(w, r)
	case "POST":
		handlePostExpense(w, r)
	case "DELETE":
		handleDeleteExpense(w, r)
	case "PATCH":
		handlePatchExpense(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
}

func handleGetExpenses(w http.ResponseWriter, r *http.Request) {
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

	jsonData, err := json.Marshal(expensesData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error converting expenses to JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func handlePostExpense(w http.ResponseWriter, r *http.Request) {
	var newExpense db.InsertExpenseDTO
	if err := parseAndValidateJsonBody(r.Body, &newExpense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if dbErr := services.GetExpensesService().CreateExpense(newExpense); dbErr != nil {
		handleError(w, dbErr, "Error creating expense", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func handleDeleteExpense(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if dbErr := services.GetExpensesService().DeleteExpense(id); dbErr != nil {
		handleError(w, dbErr, "Error deleting expense", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func handlePatchExpense(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var updateExpense db.UpdateExpenseDTO
	if err := parseAndValidateJsonBody(r.Body, &updateExpense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if dbErr := services.GetExpensesService().UpdateExpense(id, updateExpense); dbErr != nil {
		handleError(w, dbErr, "Error updating expense", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
