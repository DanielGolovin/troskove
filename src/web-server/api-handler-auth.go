package web_server

import (
	"errors"
	"log"
	"net/http"
	"time"
	"troskove/services"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Auth handler")
	log.Print(r.URL.Path)

	switch r.URL.Path {
	case "/api/login":
		handleLogin(w, r)
	case "/api/logout":
		handleLogout(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	pageHandlerIndex(w, r)
}

type LoginDTO struct {
	Token string
}

func (c *LoginDTO) Validate() error {
	if c.Token == "" {
		return errors.New("Token is required")
	}
	return nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials LoginDTO
	if err := parseAndValidateJsonBody(r.Body, &credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := services.GetAuthService().VerifyToken(credentials.Token)

	if err != nil {
		handleError(w, err, "Error logging in", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    credentials.Token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour * 30 * 12),
	}

	http.SetCookie(w, cookie)

	updateRequestCookie(r, cookie)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	updateRequestCookie(r, &http.Cookie{
		Name:  "token",
		Value: "",
	})
}
