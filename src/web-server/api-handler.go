package web_server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Validatable interface {
	Validate() error
}

func handleError(w http.ResponseWriter, err error, message string, code int) {
	log.Printf("Error: %s: %v", message, err)
	http.Error(w, message, code)
}

func parseAndValidateJsonBody(bodyReader io.ReadCloser, buffer Validatable) error {
	body, err := io.ReadAll(bodyReader)

	if err != nil {
		return errors.New("error reading request body")
	}

	if err := json.Unmarshal(body, buffer); err != nil {
		return errors.New("error parsing JSON body")
	}

	if err := buffer.Validate(); err != nil {
		return err
	}

	return nil
}
