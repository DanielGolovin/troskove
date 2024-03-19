package telegram_bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCallbackQuery(bot *tgbotapi.BotAPI, api IApi, callbackQuery tgbotapi.CallbackQuery) {
	key := callbackQuery.Data

	log.Println(key)

	if len(key) == 0 {
		editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Invalid callback data")
		return
	}

	transaction, ok := Callbackdatamap[key]

	if !ok {
		editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Callback query expired")
		return
	}

	delete(Callbackdatamap, key)

	err := api.addTransaction(transaction)

	if err != nil {
		log.Println(err)
		editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Failed to add expense")
		return
	}

	editMsg := fmt.Sprintf("Expense added: %f", transaction.Amount)
	editMessage(bot, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, editMsg)
}
