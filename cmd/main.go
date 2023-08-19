package main

import (
	"github.com/daddydemir/assistant/internal/log"
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/rabbit"
)

func main() {
	log.InitLogger()
	log.Infoln("Logger started.")

	config.Initialize()

	go rabbit.MessageQueue()
	rabbit.ImageQueue()

	for {

	}
}
