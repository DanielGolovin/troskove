package main

import (
	telegram_bot "troskove/bot"
	financial_management_infrastructure "troskove/finantial-management/infrastructure"
	webserverv "troskove/web-server"

	_ "modernc.org/sqlite"
)

func main() {
	financial_management_infrastructure.SetupDB()

	go telegram_bot.SetupBot()

	webserverv.Serve()
}
