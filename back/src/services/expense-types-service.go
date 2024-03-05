package services

import (
	"database/sql"
	"troskove/db"
)

var expenseTypesServiceInstance IExpenseTypesService

func GetExpenseTypeService() IExpenseTypesService {
	if expenseTypesServiceInstance == nil {
		expenseTypesServiceInstance = &ExpenseTypesService{
			DB: db.GetDBConnection(),
		}
	}

	return expenseTypesServiceInstance
}

type ExpenseType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExpenseTypesService struct {
	DB *sql.DB
}

type IExpenseTypesService interface {
	GetExpenseTypes() ([]ExpenseType, error)
	CreateExpenseType(expenseType db.InsertExpenseTypeDTO) error
	UpdateExpenseType(id string, expenseType db.UpdateExpenseTypeDTO) error
	DeleteExpenseType(id string) error
}

func (ets *ExpenseTypesService) GetExpenseTypes() ([]ExpenseType, error) {
	rows, err := db.GetExpenseTypesQuery(ets.DB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	expenseTypes, err := parseExpenseTypes(rows)

	if err != nil {
		return nil, err
	}

	return expenseTypes, nil
}

func (ets *ExpenseTypesService) CreateExpenseType(expenseType db.InsertExpenseTypeDTO) error {
	_, err := db.InsertExpenseType(ets.DB, expenseType)
	return err
}

func (ets *ExpenseTypesService) UpdateExpenseType(id string, expenseType db.UpdateExpenseTypeDTO) error {
	_, err := db.UpdateExpenseType(ets.DB, id, expenseType)
	return err
}

func (ets *ExpenseTypesService) DeleteExpenseType(id string) error {
	_, err := db.DeleteExpenseType(ets.DB, id)
	return err
}

func parseExpenseTypes(rows *sql.Rows) ([]ExpenseType, error) {
	var expenseTypes []ExpenseType

	for rows.Next() {
		var et ExpenseType
		if err := rows.Scan(&et.ID, &et.Name); err != nil {
			return nil, err
		}
		expenseTypes = append(expenseTypes, et)
	}

	return expenseTypes, nil
}
