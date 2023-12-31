package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Channel *amqp.Channel

func rabbitMQ() {

	conn, err := amqp.Dial(Get("RABBIT_MQ_URL"))

	if err != nil {
		log.Fatal("RabbitMQ connection was failed : ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}

	Channel = ch
}
