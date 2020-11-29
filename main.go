package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const BotToken = "BOT_TOKEN"

func main() {
	loadEnvVariables()

	botToken := os.Getenv(BotToken)
	botApi := createBotApiClient(botToken)

	c := make(chan string, 1)

	c <- botApi.Token

	for {
		go start(c)
	}
}

func loadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("oh my god")
	}
}

func createBotApiClient(botToken string) *tgbotapi.BotAPI {
	botApi, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		log.Fatal("oh my god")
	}

	return botApi
}

func start(c <-chan string) {
	fmt.Println(<-c)
}
