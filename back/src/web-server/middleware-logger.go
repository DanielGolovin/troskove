package web_server

import (
	"log"
	"net/http"
	"time"
)

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(startTime)
		log.Printf("Request handled: %s %s in %v", r.Method, r.URL.Path, duration)
	})
}
