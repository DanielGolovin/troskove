package web_server

import (
	"fmt"
	"log"
	"net/http"
)

func Setup() {
	mux := http.NewServeMux()
	port := getServerPort()

	registerRoute(mux, "/api/login", authHandler)
	registerRoute(mux, "/api/logout", authHandler)

	registerRoute(mux, "/api/expense-type", chainMiddleware(middlewareAuth)(http.HandlerFunc(expenseTypesHandler)).ServeHTTP)
	registerRoute(mux, "/api/expense", chainMiddleware(middlewareAuth)(http.HandlerFunc(expensesHandler)).ServeHTTP)

	registerRoute(mux, "/", pageHandlerIndex)

	log.Printf("Server started on port %s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
