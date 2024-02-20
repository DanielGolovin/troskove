package telegram_bot

import (
	"log"
	"regexp"
	"strconv"
	"troskove/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func getExpenseAmountFromMessage(message string) int {
	re := regexp.MustCompile("[0-9]+")

	match := re.FindString(message)

	amount, err := strconv.Atoi(match)

	if err != nil {
		amount = 0
	}

	return amount
}

var Callbackdatamap = make(map[string]ExpenseCallbackData)

func getExpenseTypeMenu(message tgbotapi.Message) tgbotapi.InlineKeyboardMarkup {
	expenseTypes, err := services.GetExpenseTypeService().GetExpenseTypes()

	if err != nil {
		log.Println(err.Error())

		return tgbotapi.NewInlineKeyboardMarkup()
	}

	var buttons [][]tgbotapi.InlineKeyboardButton

	counter := 0
	row := []tgbotapi.InlineKeyboardButton{}
	for _, expenseType := range expenseTypes {
		ecd := ExpenseCallbackData{
			ExpenseTypeId: expenseType.ID,
			Amount:        getExpenseAmountFromMessage(message.Text),
		}

		key := uuid.New().String()

		// because of 64 char limit for callback data
		Callbackdatamap[key] = ecd

		row = append(row, tgbotapi.NewInlineKeyboardButtonData(expenseType.Name, key))
		counter++
		if counter == 3 {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(row...))
			row = []tgbotapi.InlineKeyboardButton{}
			counter = 0
		}
	}

	if len(row) > 0 {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(row...))
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		buttons...,
	)

	return keyboard
}

func sendExpenseTypeMenu(bot *tgbotapi.BotAPI, message tgbotapi.Message) {
	keyboard := getExpenseTypeMenu(message)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose an option:")
	msg.ReplyMarkup = keyboard
	msg.ReplyToMessageID = message.MessageID

	bot.Send(msg)
}
