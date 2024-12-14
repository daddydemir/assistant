package notifier

import (
	"github.com/daddydemir/assistant/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
)

var api *tgbotapi.BotAPI

func GetApi() *tgbotapi.BotAPI {
	var err error

	if api == nil {
		api, err = tgbotapi.NewBotAPI(config.Get("TELEGRAM_TOKEN"))
		if err != nil {
			slog.Error("GetApi", "error", err, "token", config.Get("TELEGRAM_TOKEN"), "path", config.Get("TELEGRAM_PATH"))
		}
	}

	return api
}
