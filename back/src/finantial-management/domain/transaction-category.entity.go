package financial_management_domain

import (
	"fmt"

	"github.com/google/uuid"
)

type TransactionCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewTransactionCategory(name string) (*TransactionCategory, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return &TransactionCategory{
		ID:   uuid.New().String(),
		Name: name,
	}, nil
}
