package broker

type Broker interface {
	ReadMessage(queueName string) (<-chan any, error)
}
