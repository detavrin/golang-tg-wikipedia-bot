package main

import (
	"log"
	"os"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const language = "en"

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	// Turn off Telegram Bot API debug info
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil ||
			reflect.TypeOf(update.Message.Text).Kind() != reflect.String ||
			update.Message.Text == "" {
			continue
		}

		switch update.Message.Text {
		case "/help":

			text := "This bot searches Wikipedia articles for you. Just " +
				"enter the topic of interest and he will select several options for you.✨✨✨"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)

		case "/greetings":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🤗")
			bot.Send(msg)

		default:
			url, _ := URLEncoded(update.Message.Text)
			request := "https://" + language +
				".wikipedia.org/w/api.php?action=opensearch&search=" + url +
				"&limit=3&origin=*&format=json"
			message := WikipediaAPI(request)

			//Loop throug message slice
			for _, val := range message {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, val)
				bot.Send(msg)
			}
		}
	}
}
