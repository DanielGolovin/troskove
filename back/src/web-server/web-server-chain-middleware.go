package web_server

import (
	"net/http"
)

func chainMiddleware(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		return handler
	}
}
