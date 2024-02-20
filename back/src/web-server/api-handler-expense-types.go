package web_server

import (
	"encoding/json"
	"net/http"
	"troskove/db"
	"troskove/services"
)

func expenseTypesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetExpenseTypes(w, r)
	case "POST":
		handlePostexpenseType(w, r)
	case "DELETE":
		handleDeleteExpenseType(w, r)
	case "PATCH":
		handlePatchExpenseType(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
}

func handleGetExpenseTypes(w http.ResponseWriter, r *http.Request) {
	expenseTypes, dbErr := services.GetExpenseTypeService().GetExpenseTypes()

	if dbErr != nil {
		handleError(w, dbErr, "Error getting expense types", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(expenseTypes)

	if err != nil {
		handleError(w, err, "Error converting expense types to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func handlePostexpenseType(w http.ResponseWriter, r *http.Request) {
	newExpenseType := db.InsertExpenseTypeDTO{}
	if err := parseAndValidateJsonBody(r.Body, &newExpenseType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbErr := services.GetExpenseTypeService().CreateExpenseType(newExpenseType)

	if dbErr != nil {
		handleError(w, dbErr, "Error creating expense type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleDeleteExpenseType(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	dbErr := services.GetExpenseTypeService().DeleteExpenseType(id)

	if dbErr != nil {
		handleError(w, dbErr, "Error deleting expense type", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func handlePatchExpenseType(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	updateExpenseType := db.UpdateExpenseTypeDTO{}
	if err := parseAndValidateJsonBody(r.Body, &updateExpenseType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbErr := services.GetExpenseTypeService().UpdateExpenseType(id, updateExpenseType)

	if dbErr != nil {
		handleError(w, dbErr, "Error updating expense type", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
