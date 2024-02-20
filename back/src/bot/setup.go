package telegram_bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupBot() *tgbotapi.BotAPI {
	botToken := getBotToken()
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Setting up bot update channel...")
	updatesChan := setupUpdateChannel(bot)
	log.Println("Bot update channel is set...")

	log.Println("Bot is running...")
	for update := range updatesChan {
		handleUpdate(bot, update)
	}

	// bot.Debug = true

	return bot
}

func getBotToken() string {
	botToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")

	if botToken == "" {
		log.Fatalln("TELEGRAM_BOT_API_TOKEN is not set")
	}

	return botToken
}

func setupUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	return updateConfig
}

func setupUpdateChannel(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	updateConfig := setupUpdateConfig()
	updates := bot.GetUpdatesChan(updateConfig)

	return updates
}
