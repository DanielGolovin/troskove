package webserverv

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

	mux.HandleFunc("GET /transactions", chain(transactionHandler.GetTransactions, middlewareLogger))
	mux.HandleFunc("POST /transactions", chain(transactionHandler.CreateTransaction, middlewareLogger))

	// transaction categories
	transactionCategoryService := financial_management_application.NewTransactionCategoryService(transactionCategoryRepository)
	transactionCategoryHandler := financial_management_presentation_rest.NewTransactionCategorynRestHandler(transactionCategoryService)

	mux.HandleFunc("POST /transaction-categories", chain(transactionCategoryHandler.CreateTransactionCategory, middlewareLogger))
	mux.HandleFunc("GET /transaction-categories", chain(transactionCategoryHandler.GetTransactionCategories, middlewareLogger))
	mux.HandleFunc("PUT /transaction-categories/{id}", chain(transactionCategoryHandler.UpdateTransactionCategory, middlewareLogger))
	mux.HandleFunc("DELETE /transaction-categories/{id}", chain(transactionCategoryHandler.DeleteTransactionCategory, middlewareLogger))

	log.Println("Listening on port ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
