package financial_management_domain

type ITransactionRepository interface {
	Create(transaction Transaction) error
	FindAll(options FindAllOptions) ([]Transaction, error)
	Delete(id string) error
	Update(id string, transaction Transaction) error
}

type FindAllOptions struct {
	OrderBy  string
	OrderDir string
	Limit    int
	Offset   int
	Filters  FindAllFilters
}

type FindAllFilters struct {
	CategoryID  string
	StartDate   string
	EndDate     string
	BiggerThan  string
	SmallerThan string
}
