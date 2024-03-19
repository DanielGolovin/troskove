package financial_management_application

func (service *TransactionService) DeleteTransaction(id string) error {
	return service.transactionRepository.Delete(id)
}
