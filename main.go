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
	err := godotenv.Load(".env")
	c := make(chan string, 1)

	if err != nil {
		log.Fatal("oh my god")
	}

	botToken := os.Getenv(BotToken)

	_, err = tgbotapi.NewBotAPI(botToken)

	if err != nil {
		log.Fatal("oh my god")
	}

	c <- botToken

	for {
		go start(c)
	}
}

func start(c <-chan string) {
	fmt.Println(<-c)
}
