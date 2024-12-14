package service

import (
	"github.com/daddydemir/assistant/pkg/broker"
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/notifier"
	"log/slog"
)

func SendImage(broker broker.Broker, notifier notifier.Notifier) {
	messages, err := broker.ReadMessage(config.Get("QUEUE_NAME_2"))
	if err != nil {
		slog.Error("SendImage", "error", err)
	}

	for msg := range messages {
		switch m := msg.(type) {
		case []byte:
			slog.Info("Received message:", "message", m)
			notifier.NotifyImage(string(m))
		default:
			notifier.NotifyMessage("Sanirim bir hata olustu! {image}")
			slog.Error("Unknown message:", "message", m)
		}
	}
}

func SendMessage(broker broker.Broker, notifier notifier.Notifier) {
	messages, err := broker.ReadMessage(config.Get("QUEUE_NAME_1"))
	if err != nil {
		slog.Error("SendMessage", "error", err)
	}

	for msg := range messages {
		switch m := msg.(type) {
		case []byte:
			slog.Info("Received message:", "message", m)
			notifier.NotifyMessage(string(m))
		default:
			notifier.NotifyMessage("Sanirim bir hata olustu!")
			slog.Error("Unknown message:", "message", m)
		}
	}
}
