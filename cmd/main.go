package main

import (
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/rabbit"
)

func main() {

	config.Initialize()

	go rabbit.MessageQueue()
	rabbit.ImageQueue()

	for {

	}
}
