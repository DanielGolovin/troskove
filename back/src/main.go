package main

import (
	telegram_bot "troskove/bot"
	financial_management_infrastructure "troskove/finantial-management/infrastructure"
	webserverv2 "troskove/web-server-v2"

	_ "modernc.org/sqlite"
)

func main() {
	financial_management_infrastructure.SetupDB()

	go telegram_bot.SetupBot()

	webserverv2.Serve()
}
