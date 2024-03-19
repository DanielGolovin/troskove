package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUpdate(bot *tgbotapi.BotAPI, api IApi, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		handleCallbackQuery(bot, api, *update.CallbackQuery)
	} else if update.Message != nil {
		sendExpenseTypeMenu(bot, api, *update.Message)
	}
}
