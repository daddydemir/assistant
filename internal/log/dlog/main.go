package dlog

import (
	"github.com/daddydemir/assistant/pkg/config"
	"log/slog"
	"os"
	"time"
)

var logger *slog.Logger

func init() {
	folder := config.Get("LOG_FOLDER")
	today := time.Now().Format("2006-01-02")

	file := folder + today + "-json.log"
	openFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	jsonHandler := slog.NewJSONHandler(openFile, &slog.HandlerOptions{
		AddSource: true,
	})
	logger = slog.New(jsonHandler)
	slog.SetDefault(logger)
}
