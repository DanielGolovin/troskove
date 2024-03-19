package financial_management_application

import financial_management_domain "troskove/finantial-management/domain"

type UpdateTransactionDTO struct {
	ID         string  `json:"id"`
	Amount     float64 `json:"amount"`
	Date       string  `json:"date"`
	CategoryID string  `json:"categoryId"`
}

func (service *TransactionService) UpdateTransaction(data UpdateTransactionDTO) error {
	category, err := service.transactionCategoryRepository.FindById(data.CategoryID)

	if err != nil {
		return err
	}

	newTransactionData := financial_management_domain.NewTransactionDTO{
		Amount:   data.Amount,
		Date:     data.Date,
		Category: *category,
	}

	transactionData, err := financial_management_domain.NewTransaction(newTransactionData)

	if err != nil {
		return err
	}

	return service.transactionRepository.Update(data.CategoryID, *transactionData)
}
