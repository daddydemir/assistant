package broker

import (
	"fmt"
	"github.com/daddydemir/assistant/pkg/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

var channel *amqp.Channel

func StartRabbitmq() {
	GetChannel()
}

func GetChannel() *amqp.Channel {
	if channel == nil || channel.IsClosed() {
		dial, err := amqp.Dial(config.Get("RABBIT_MQ_URL"))
		if err != nil {
			fmt.Println("Failed to connect to RabbitMQ", err)
		}

		channel, err = dial.Channel()
		if err != nil {
			fmt.Println("Failed to open a channel", err)
		}
	}

	return channel
}
