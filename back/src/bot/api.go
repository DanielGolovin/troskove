package telegram_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IApi interface {
	addTransaction(data CreateTransactionDTO) error
	getTransactionTypes() ([]TransactionCategory, error)
}

type CreateTransactionDTO struct {
	Amount     float64 `json:"amount"`
	Date       string  `json:"date"`
	CategoryID string  `json:"categoryId"`
}

type TransactionCategory struct {
	ID   string
	Name string
}

type Api struct {
}

func NewApi() IApi {
	return &Api{}
}

const url = "http://localhost:8080"

func (a *Api) addTransaction(data CreateTransactionDTO) error {
	addUrl := url + "/transactions"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}

	resp, err := http.Post(addUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return fmt.Errorf("API error: %s", resp.Status)
		}

		return fmt.Errorf("API error: %s", string(body))
	}

	return nil
}

func (a *Api) getTransactionTypes() ([]TransactionCategory, error) {
	categoriesUrl := url + "/transaction-categories"

	resp, err := http.Get(categoriesUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("API error: %s", resp.Status)
		}

		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var categories []TransactionCategory
	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return categories, nil
}
