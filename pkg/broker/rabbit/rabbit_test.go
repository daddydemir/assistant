package rabbit

import (
	"fmt"
	"testing"
)

func TestConsumer_ReadMessage(t *testing.T) {
	consumer := Consumer{}
	messages, err := consumer.ReadMessage("queueName")
	if err != nil {
		t.Error(err)
	}

	for msg := range messages {
		switch v := msg.(type) {
		case []byte:
			fmt.Printf("Received RabbitMQ message: %s\n", string(v))
		default:
			fmt.Println("Unknown message type")
		}
	}

}
