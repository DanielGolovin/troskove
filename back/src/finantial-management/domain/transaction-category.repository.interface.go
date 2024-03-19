package financial_management_domain

type ITransactionCategoryRepository interface {
	Create(transactionCategory TransactionCategory) error
	FindById(id string) (*TransactionCategory, error)
	FindAll() ([]TransactionCategory, error)
	Update(id string, transactionCategory TransactionCategory) error
	Delete(id string) error
}
