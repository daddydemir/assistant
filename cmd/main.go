package main

import (
	"github.com/daddydemir/assistant/internal/log"
	_ "github.com/daddydemir/assistant/internal/log/dlog"
	"github.com/daddydemir/assistant/pkg/broker/rabbit"
	"github.com/daddydemir/assistant/pkg/config/broker"
	"github.com/daddydemir/assistant/pkg/notifier/telegram"
	"github.com/daddydemir/assistant/pkg/service"
)

func main() {
	log.InitLogger()
	log.Infoln("Logger started.")

	broker.StartRabbitmq()

	consumer := &rabbit.Consumer{}
	notifier := &telegram.TelegramNotifier{}

	go service.SendMessage(consumer, notifier)
	service.SendImage(consumer, notifier)

	select {}
}
