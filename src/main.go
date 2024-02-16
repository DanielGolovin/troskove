package main

import (
	"log"
	telegram_bot "troskove/bot"
	"troskove/db"
	web_server "troskove/web-server"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Setting up envs...")
	setupEnvs()
	log.Println("Envs are set...")

	log.Println("Setting up database...")
	db.SetupDB()
	log.Println("Database is set...")

	log.Println("Setting up web server...")
	go web_server.Setup()
	log.Println("Web server is running...")

	log.Println("Setting up bot...")
	telegram_bot.SetupBot()
	log.Println("Execution is finished...")
}

func setupEnvs() {
	log.Println("Loading .env file...")

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found...")
	} else {
		log.Println(".env file loaded...")
	}
}
