package telegram_bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ExpenseCallbackData struct {
	Amount        int
	ExpenseTypeId string
}

type ExpenseTypeCallbackData struct {
	Name string
}

func handleUnauthorized(bot *tgbotapi.BotAPI, chatID int64, userId int64) {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("You are not authorized to use this bot. Ask the admin to add you. Your user ID is: %d", userId))
	bot.Send(msg)
}

func getAllowedUsers() []int64 {
	userIdString := os.Getenv("ALLOWED_TELEGRAM_USER_IDS")

	userIdStrings := strings.Split(userIdString, ",")

	userIds := make([]int64, len(userIdStrings))

	for _, userIdString := range userIdStrings {
		userId, err := strconv.ParseInt(userIdString, 10, 64)
		if err != nil {
			log.Fatalf("Error parsing user ID: %s", userIdString)
		}

		userIds = append(userIds, userId)
	}

	return userIds
}

func isAllowedToUseBot(bot *tgbotapi.BotAPI, chatID int64, userId int64) bool {
	allowedUsers := getAllowedUsers()

	for _, allowedUser := range allowedUsers {
		if userId == allowedUser {
			return true
		}
	}

	return false
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if isAllowed := isAllowedToUseBot(bot, update.Message.Chat.ID, update.Message.From.ID); !isAllowed {
		handleUnauthorized(bot, update.Message.Chat.ID, update.Message.From.ID)
		return
	}

	if update.CallbackQuery != nil {
		handleCallbackQuery(bot, *update.CallbackQuery)
	} else if update.Message != nil {
		if update.Message.IsCommand() {
			sendToken(bot, update.Message.Chat.ID)
		} else {
			sendExpenseTypeMenu(bot, *update.Message)
		}
	}
}
