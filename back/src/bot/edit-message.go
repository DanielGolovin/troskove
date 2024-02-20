package telegram_bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func editMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int, text string) error {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, text)

	_, err := bot.Send(msg)

	return err
}
