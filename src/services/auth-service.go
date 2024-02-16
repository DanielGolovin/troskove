package services

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var authServiceInstance Authenticator

func GetAuthService() Authenticator {
	if authServiceInstance == nil {
		authServiceInstance = &AuthService{
			secretKey: getSecretKey(),
		}
	}

	return authServiceInstance
}

type Authenticator interface {
	CreateToken() (string, error)
	VerifyToken(tokenString string) (bool, error)
	VerifyRequest(r *http.Request) (bool, error)
}

type AuthService struct {
	secretKey string
}

func (as *AuthService) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})

	tokenString, err := token.SignedString([]byte(as.secretKey))
	return tokenString, err
}

func (as *AuthService) VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.secretKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (as *AuthService) VerifyRequest(r *http.Request) (bool, error) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		log.Println("Error getting token from cookie")
		return false, err
	}

	token := tokenCookie.Value

	if token == "" {
		return false, errors.New("Token is empty")
	}

	isValid, err := as.VerifyToken(token)
	return isValid, err
}

func getSecretKey() string {
	key := os.Getenv("JWT_SECRET_KEY")

	if key == "" {
		log.Fatalln("JWT_SECRET_KEY is not set")
	}

	return key
}
