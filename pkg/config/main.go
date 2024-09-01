package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var path string

func init() {
	path = os.Getenv("ENV_PATH")
	if path == "" {
		path = "../configs/prod.env"
	}
}

func Get(key string) string {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
