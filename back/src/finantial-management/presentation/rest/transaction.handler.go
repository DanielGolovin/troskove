package financial_management_presentation_rest

import (
	"encoding/json"
	"net/http"
	financial_management_application "troskove/finantial-management/application"
	financial_management_domain "troskove/finantial-management/domain"
)

type TransactionRestHandler struct {
	TransactionService financial_management_application.ITransactionService
}

type ITransactionRestHandler interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	UpdateTransaction(w http.ResponseWriter, r *http.Request)
	DeleteTransaction(w http.ResponseWriter, r *http.Request)
}

func NewTransactionRestHandler(transactionService financial_management_application.ITransactionService) ITransactionRestHandler {
	return &TransactionRestHandler{
		TransactionService: transactionService,
	}
}

func (handler *TransactionRestHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction financial_management_application.CreateTransactionDTO
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.TransactionService.Create(transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *TransactionRestHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	if handler.TransactionService == nil {
		http.Error(w, "TransactionService is nil", http.StatusInternalServerError)
		return
	}

	filters := financial_management_domain.FindAllFilters{
		CategoryID:  r.URL.Query().Get("categoryId"),
		StartDate:   r.URL.Query().Get("from"),
		EndDate:     r.URL.Query().Get("to"),
		BiggerThan:  r.URL.Query().Get("biggerThan"),
		SmallerThan: r.URL.Query().Get("smallerThan"),
	}

	options := financial_management_domain.FindAllOptions{
		OrderBy:  r.URL.Query().Get("orderBy"),
		OrderDir: r.URL.Query().Get("orderDir"),
		Limit:    25,
		Offset:   0,
		Filters:  filters,
	}

	transactions, err := handler.TransactionService.GetAll(options)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (handler *TransactionRestHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction financial_management_application.UpdateTransactionDTO
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.TransactionService.UpdateTransaction(transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (handler *TransactionRestHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("id")

	err := handler.TransactionService.DeleteTransaction(transactionID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
