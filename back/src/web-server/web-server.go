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

	registerRoute(mux, "/api/expense-type", chainMiddleware()(http.HandlerFunc(expenseTypesHandler)).ServeHTTP)
	registerRoute(mux, "/api/expense", chainMiddleware()(http.HandlerFunc(expensesHandler)).ServeHTTP)

	log.Printf("Server started on port %s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
