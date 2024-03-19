package webserverv2

import (
	"fmt"
	"log"
	"net/http"
	financial_management_application "troskove/finantial-management/application"
	financial_management_infrastructure "troskove/finantial-management/infrastructure"
	financial_management_presentation_rest "troskove/finantial-management/presentation/rest"
)

func Serve() {
	mux := http.NewServeMux()
	port := "8080"

	dbConnection := financial_management_infrastructure.GetDBConnection()

	transactionsRepository := financial_management_infrastructure.NewTransactionRepository(dbConnection)
	transactionCategoryRepository := financial_management_infrastructure.NewTransactionCategoryRepository(dbConnection)

	// transactions
	transactionService := financial_management_application.NewTransactionService(transactionsRepository, transactionCategoryRepository)
	transactionHandler := financial_management_presentation_rest.NewTransactionRestHandler(transactionService)

	mux.HandleFunc("GET /transactions", transactionHandler.GetTransactions)
	mux.HandleFunc("POST /transactions", transactionHandler.CreateTransaction)

	// transaction categories
	transactionCategoryService := financial_management_application.NewTransactionCategoryService(transactionCategoryRepository)
	transactionCategoryHandler := financial_management_presentation_rest.NewTransactionCategorynRestHandler(transactionCategoryService)

	mux.HandleFunc("POST /transaction-categories", transactionCategoryHandler.CreateTransactionCategory)
	mux.HandleFunc("GET /transaction-categories", transactionCategoryHandler.GetTransactionCategories)
	mux.HandleFunc("PUT /transaction-categories/{id}", transactionCategoryHandler.UpdateTransactionCategory)
	mux.HandleFunc("DELETE /transaction-categories/{id}", transactionCategoryHandler.DeleteTransactionCategory)

	log.Println("Listening on port ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
