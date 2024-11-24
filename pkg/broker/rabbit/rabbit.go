package rabbit

import (
	"context"
	"fmt"
	"github.com/daddydemir/assistant/pkg/config/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type Consumer struct{}

func (c *Consumer) ReadMessage(queueName string) (<-chan any, error) {

	channel := broker.GetChannel()

	outChan := make(chan any)

	messages, err := channel.Consume(
		getQueue(queueName).Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("Error consuming message", err)
		return nil, err
	}

	go func() {
		for msg := range messages {
			outChan <- msg.Body
		}
		close(outChan)
	}()

	return outChan, nil
}

func getQueue(queueName string) amqp.Queue {

	channel := broker.GetChannel()

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to declare a queue:", err)
	}

	return queue
}

func (c *Consumer) WriteMessage(queueName string, message string) error {
	channel := broker.GetChannel()
	ctx := context.Background()

	err := channel.PublishWithContext(ctx, "", getQueue(queueName).Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	if err != nil {
		slog.Error("Failed to publish a message", err)
		return err
	}
	return nil
}
