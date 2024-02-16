package web_server

import (
	"net/http"
)

func registerRoute(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	handlerFunc := http.HandlerFunc(handler)
	wrappedHandler := chainMiddleware(middlewareLogger)(handlerFunc)
	mux.HandleFunc(path, wrappedHandler.ServeHTTP)
}
