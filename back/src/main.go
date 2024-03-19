package main

import (
	financial_management_infrastructure "troskove/finantial-management/infrastructure"
	webserverv2 "troskove/web-server-v2"

	_ "modernc.org/sqlite"
)

func main() {
	financial_management_infrastructure.SetupDB()

	webserverv2.Serve()
}
