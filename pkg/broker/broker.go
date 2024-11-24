package broker

import "github.com/daddydemir/assistant/pkg/broker/rabbit"

type Broker interface {
	ReadMessage(queueName string) (<-chan any, error)
	WriteMessage(queueName string, message string) error
}

func GetActiveBroker() Broker {
	return &rabbit.Consumer{}
}
