package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zenKmi/epubify-tools/cmd/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	BASE_URL := os.Getenv("BASE_URL")
	PORT := os.Getenv("PORT")

	server := api.SetupRouter()
	server.Run(BASE_URL + ":" + PORT)
}
