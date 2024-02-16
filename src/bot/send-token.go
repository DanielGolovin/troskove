package telegram_bot

import (
	"troskove/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendToken(bot *tgbotapi.BotAPI, chatID int64) {
	token, err := services.GetAuthService().CreateToken()

	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Error creating token"))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, token))
}
