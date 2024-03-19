package webserverv

import "net/http"

func chain(f http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}
