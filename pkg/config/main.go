package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Initialize() {
	rabbitMQ()
	log.Println("RabbitMQ connection is start.")
	telegram()
	log.Println("Telegram bot is active.")
}

func Get(key string) string {
	err := godotenv.Load("C:\\Users\\demir\\dev\\assistant\\configs\\prod.env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
