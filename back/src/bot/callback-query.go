package telegram_bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery tgbotapi.CallbackQuery) {
	key := callbackQuery.Data

	log.Println(key)

	if len(key) == 0 {
		editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Invalid callback data")
		return
	}

	expense, ok := Callbackdatamap[key]

	if !ok {
		editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Callback query expired")
		return
	}

	delete(Callbackdatamap, key)
	addExpense(expense.ExpenseTypeId, expense.Amount)

	editMsg := fmt.Sprintf("Expense added: %d", expense.Amount)
	editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, editMsg)
}
