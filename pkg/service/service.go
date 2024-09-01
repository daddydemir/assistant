package service

import (
	"fmt"
	"github.com/daddydemir/assistant/pkg/broker"
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/notifier"
)

func SendImage(broker broker.Broker, notifier notifier.Notifier) {
	messages, err := broker.ReadMessage(config.Get("QUEUE_NAME_2"))
	if err != nil {
		fmt.Println(err)
	}

	for msg := range messages {
		switch m := msg.(type) {
		case []byte:
			notifier.NotifyImage(string(m))
		default:
			notifier.NotifyMessage("Sanirim bir hata olustu! {image}")
			fmt.Println("Unknown message:", m)
		}
	}
}

func SendMessage(broker broker.Broker, notifier notifier.Notifier) {
	messages, err := broker.ReadMessage(config.Get("QUEUE_NAME_1"))
	if err != nil {
		fmt.Println(err)
	}

	for msg := range messages {
		switch m := msg.(type) {
		case []byte:
			notifier.NotifyMessage(string(m))
		default:
			notifier.NotifyMessage("Sanirim bir hata olustu!")
			fmt.Println("Unknown message:", m)
		}
	}
}
