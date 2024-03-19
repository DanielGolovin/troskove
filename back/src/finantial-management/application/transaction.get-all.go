package financial_management_application

import financial_management_domain "troskove/finantial-management/domain"

func (service *TransactionService) GetAll(options financial_management_domain.FindAllOptions) ([]financial_management_domain.Transaction, error) {

	return service.transactionRepository.FindAll(options)
}
