package web_server

import (
	"log"
	"net/http"
	"troskove/services"
)

func middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			log.Println("No token provided")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		isValid, err := services.GetAuthService().VerifyRequest(r)

		if err != nil {
			log.Println("Error verifying token: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !isValid {
			log.Println("Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
