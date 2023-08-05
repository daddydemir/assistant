package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Initialize() {
	rabbitMQ()
	println("RabbitMQ connection is start.")
	telegram()
	println("Telegram bot is active.")
}

func Get(key string) string {
	err := godotenv.Load("C:\\Users\\demir\\dev\\assistant\\configs\\prod.env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
