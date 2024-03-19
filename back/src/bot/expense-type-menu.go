package telegram_bot

import (
	"log"
	"regexp"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func getExpenseAmountFromMessage(message string) float64 {
	re := regexp.MustCompile("[0-9]+")

	match := re.FindString(message)

	amount, err := strconv.ParseFloat(match, 64)

	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return amount
}

var Callbackdatamap = make(map[string]CreateTransactionDTO)

func getExpenseTypeMenu(message tgbotapi.Message, transactionTypes []TransactionCategory) tgbotapi.InlineKeyboardMarkup {
	var buttons [][]tgbotapi.InlineKeyboardButton

	unixDate := message.Date
	date := time.Unix(int64(unixDate), 0).Format(time.RFC3339)

	counter := 0
	row := []tgbotapi.InlineKeyboardButton{}
	for _, expenseType := range transactionTypes {
		ecd := CreateTransactionDTO{
			CategoryID: expenseType.ID,
			Amount:     getExpenseAmountFromMessage(message.Text),
			Date:       date,
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

func sendExpenseTypeMenu(bot *tgbotapi.BotAPI, api IApi, message tgbotapi.Message) {
	// TODO: use go routines
	transactionTypes, err := api.getTransactionTypes()

	if err != nil {
		log.Println(err.Error())
		return
	}

	keyboard := getExpenseTypeMenu(message, transactionTypes)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose an option:")
	msg.ReplyMarkup = keyboard
	msg.ReplyToMessageID = message.MessageID

	bot.Send(msg)
}
