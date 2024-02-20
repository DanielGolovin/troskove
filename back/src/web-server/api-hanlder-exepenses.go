package web_server

import (
	"net/http"
	"troskove/db"
	"troskove/services"
)

func expensesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
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

	pageHandlerIndex(w, r)
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
}
