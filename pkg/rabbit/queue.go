package rabbit

import (
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/telegram"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func getQueue(name string) amqp091.Queue {

	q, err := config.Channel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("QueueDeclare has occurred error : ", err)
	}
	return q
}

func MessageQueue() {

	messages, err := config.Channel.Consume(
		getQueue(config.Get("QUEUE_NAME_1")).Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	var forever chan struct{}

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
			telegram.SendMessage(string(d.Body))
		}
	}()

	<-forever
}

func ImageQueue() {
	messages, err := config.Channel.Consume(
		getQueue(config.Get("QUEUE_NAME_2")).Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	var images chan struct{}

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
			telegram.SendImage(string(d.Body))
		}
	}()

	<-images
}
