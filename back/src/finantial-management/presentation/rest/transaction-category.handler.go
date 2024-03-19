package financial_management_presentation_rest

import (
	"encoding/json"
	"net/http"
	financial_management_application "troskove/finantial-management/application"
)

type TransactionCategorynRestHandler struct {
	TransactionCategorynService financial_management_application.ITransactionCategoryService
}

type ITransactionCategorynRestHandler interface {
	CreateTransactionCategory(w http.ResponseWriter, r *http.Request)
	GetTransactionCategories(w http.ResponseWriter, r *http.Request)
	DeleteTransactionCategory(w http.ResponseWriter, r *http.Request)
	UpdateTransactionCategory(w http.ResponseWriter, r *http.Request)
}

func NewTransactionCategorynRestHandler(transactionCategorynService financial_management_application.ITransactionCategoryService) ITransactionCategorynRestHandler {
	return &TransactionCategorynRestHandler{
		TransactionCategorynService: transactionCategorynService,
	}
}

type CreateTransactionCategoryDTO struct {
	Name string
}

func (handler *TransactionCategorynRestHandler) CreateTransactionCategory(w http.ResponseWriter, r *http.Request) {
	var transactionCategory CreateTransactionCategoryDTO
	err := json.NewDecoder(r.Body).Decode(&transactionCategory)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.TransactionCategorynService.Create(transactionCategory.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TransactionCategorynRestHandler) GetTransactionCategories(w http.ResponseWriter, r *http.Request) {
	transactionCategories, err := handler.TransactionCategorynService.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactionCategories)
}

func (handler *TransactionCategorynRestHandler) DeleteTransactionCategory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := handler.TransactionCategorynService.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (handler *TransactionCategorynRestHandler) UpdateTransactionCategory(w http.ResponseWriter, r *http.Request) {
	var transactionCategory financial_management_application.UpdateTransactionCategoryDTO
	err := json.NewDecoder(r.Body).Decode(&transactionCategory)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if transactionCategory.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err = handler.TransactionCategorynService.Update(id, transactionCategory)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
