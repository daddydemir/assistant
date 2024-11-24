package main

import (
	"github.com/daddydemir/assistant/handler"
	_ "github.com/daddydemir/assistant/internal/log/dlog"
	"github.com/daddydemir/assistant/pkg/broker/rabbit"
	"github.com/daddydemir/assistant/pkg/config/broker"
	"github.com/daddydemir/assistant/pkg/notifier/telegram"
	"github.com/daddydemir/assistant/pkg/service"
)

func main() {
	broker.StartRabbitmq()

	consumer := &rabbit.Consumer{}
	notifier := &telegram.TelegramNotifier{}

	// http server
	go handler.StartApi()

	go service.SendMessage(consumer, notifier)
	service.SendImage(consumer, notifier)

	select {}
}
