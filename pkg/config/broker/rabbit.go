package broker

import (
	"github.com/daddydemir/assistant/pkg/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

var channel *amqp.Channel

func StartRabbitmq() {
	GetChannel()
}

func GetChannel() *amqp.Channel {
	if channel == nil || channel.IsClosed() {
		dial, err := amqp.Dial(config.Get("RABBIT_MQ_URL"))
		if err != nil {
			slog.Error("GetChannel", "error", err, "url", config.Get("RABBIT_MQ_URL"))
		}

		channel, err = dial.Channel()
		if err != nil {
			slog.Error("GetChannel", "error", err)
		}
	}

	return channel
}
