package web_server

import (
	"log"
	"os"
)

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatalln("SERVER_PORT is not set")
	}

	return port
}

func getTemplateFolderPath() string {
	path := os.Getenv("TEMPLATE_FOLDER_PATH")

	if path == "" {
		log.Fatalln("TEMPLATE_FOLDER_PATH is not set")
	}

	return path
}
