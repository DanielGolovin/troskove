package telegram_bot

import (
	"log"
	"troskove/db"
	"troskove/services"
)

func addExpense(expenseTypeId string, amount int) {
	float64Amount := float64(amount)

	expense := db.InsertExpenseDTO{
		Value:  float64Amount,
		TypeID: expenseTypeId,
	}

	err := services.GetExpensesService().CreateExpense(expense)

	if err != nil {
		log.Fatalln(err)
	}
}
