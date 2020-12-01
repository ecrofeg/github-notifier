package notifier

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Notifier struct {
	telegram *tgbotapi.BotAPI
}

func MakeNotifier(botToken string) *Notifier {
	botAPI, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		panic("oh my god")
	}

	return &Notifier{
		telegram: botAPI,
	}
}

func (n *Notifier) Listen(chatId int64, messages <-chan string) {
	for messageText := range messages {
		fmt.Println("Sending:", messageText)

		n.sendMessage(chatId, "Message: "+messageText)
	}
}

func (n *Notifier) sendMessage(chatID int64, messageText string) {
	message := tgbotapi.NewMessage(chatID, "Message: "+messageText)
	_, err := n.telegram.Send(message)

	if err != nil {
		log.Fatal("message was not send", err)
	}
}
