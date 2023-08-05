package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var Bot *tgbotapi.BotAPI

func telegram() {
	bot, err := tgbotapi.NewBotAPI(Get("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	Bot = bot
}
