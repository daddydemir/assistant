package telegram

import (
	"bytes"
	"fmt"
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/config/notifier"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strconv"
)

type TelegramNotifier struct{}

func (t *TelegramNotifier) NotifyMessage(message string) {
	chatId := config.Get("CHAT_ID")

	id, err := strconv.Atoi(chatId)
	if err != nil {
		slog.Error("NotifyMessage:Atoi", "error", err, "chatId", chatId)
	}

	api := notifier.GetApi()
	m, err := api.Send(tgbotapi.NewMessage(int64(id), message))
	if err != nil {
		slog.Error("NotifyMessage:send", "error", err, "message", message)
	}
	slog.Info("NotifyMessage:send", "message", message, "m", m)
}

func (t *TelegramNotifier) NotifyImage(imagePath string) {
	path := config.Get("IMAGE_FOLDER") + imagePath
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", config.Get("TELEGRAM_TOKEN"))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, _ := writer.CreateFormFile("photo", path)
	file, _ := ioutil.ReadFile(path)
	_, err := fileWriter.Write(file)
	if err != nil {
		slog.Error("Write", "error", err)
		return
	}
	err = writer.WriteField("chat_id", config.Get("CHAT_ID"))
	if err != nil {
		slog.Error("WriteField", "error", err)
		return
	}

	err = writer.Close()
	if err != nil {
		slog.Error("Close", "error", err)
		return
	}

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		slog.Error("Do", "error", err, "request", req)
	}
	slog.Info("Do", "request", req, "response", res)
}
