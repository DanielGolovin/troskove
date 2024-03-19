package financial_management_domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID       string              `json:"id"`
	Amount   float64             `json:"amount"`
	Date     string              `json:"date"`
	Category TransactionCategory `json:"category"`
}

type NewTransactionDTO struct {
	Amount   float64             `json:"amount"`
	Date     string              `json:"date"`
	Category TransactionCategory `json:"category"`
}

func NewTransaction(data NewTransactionDTO) (*Transaction, error) {
	if data.Category.ID == "" {
		return nil, fmt.Errorf("category ID is required")
	}

	if data.Category.Name == "" {
		return nil, fmt.Errorf("category name is required")
	}

	if data.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	if data.Date == "" {
		return nil, fmt.Errorf("date is required")
	}
	_, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format")
	}

	return &Transaction{
		ID:       uuid.New().String(),
		Amount:   data.Amount,
		Date:     data.Date,
		Category: data.Category,
	}, nil
}
