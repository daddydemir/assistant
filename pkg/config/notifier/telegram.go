package notifier

import (
	"fmt"
	"github.com/daddydemir/assistant/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var api *tgbotapi.BotAPI

func GetApi() *tgbotapi.BotAPI {
	var err error

	if api == nil {
		api, err = tgbotapi.NewBotAPI(config.Get("TELEGRAM_TOKEN"))
		if err != nil {
			fmt.Println("Failed to connect to telegram bot", err)
		}
	}

	return api
}
