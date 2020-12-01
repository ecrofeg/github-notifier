package main

import (
	"fmt"
	"log"
	"notifier/notifier"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const BotToken = "BOT_TOKEN"
const ChatId = "CHAT_ID"

func main() {
	loadEnvVariables()

	botToken := os.Getenv(BotToken)
	chatIDString := os.Getenv(ChatId)
	chatID, err := strconv.ParseInt(chatIDString, 10, 64)

	if err != nil {
		log.Fatal("wrong chat ID")
	}

	notifier := notifier.MakeNotifier(botToken)

	messages := make(chan string)

	go produceMessages(messages)

	notifier.Listen(chatID, messages)
}

func produceMessages(messages chan<- string) {
	for _, v := range []string{"1", "2", "3", "4", "5"} {
		fmt.Println("Produced:", v)
		messages <- v
	}
}

func loadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("oh my god")
	}
}
