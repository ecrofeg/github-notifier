package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const BotToken = "BOT_TOKEN"
const ChatId = "CHAT_ID"

func main() {
	loadEnvVariables()

	botToken := os.Getenv(BotToken)
	chatIdString := os.Getenv(ChatId)
	chatId, err := strconv.ParseInt(chatIdString, 10, 64)

	if err != nil {
		log.Fatal("wrong chat ID")
	}

	botApi := createBotApiClient(botToken)

	messages := make(chan string)

	go produceMessages(messages)

	listenToMessages(botApi, chatId, messages)
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

func produceMessages(messages chan<- string) {
	for _, v := range []string{"1", "2", "3", "4", "5"} {
		fmt.Println("Produced:", v)
		messages <- v
	}
}

func listenToMessages(botApi *tgbotapi.BotAPI, chatId int64, messages <-chan string) {
	for messageText := range messages {
		fmt.Println("Sending:", messageText)

		sendMessage(botApi, chatId, "Message: "+messageText)
	}
}

func sendMessage(botApi *tgbotapi.BotAPI, chatId int64, messageText string) {
	message := tgbotapi.NewMessage(chatId, "Message: "+messageText)
	_, err := botApi.Send(message)

	if err != nil {
		log.Fatal("message was not send", err)
	}
}
