package telegram

import (
	"bytes"
	"fmt"
	"github.com/daddydemir/assistant/internal/log"
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/daddydemir/assistant/pkg/config/notifier"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

type TelegramNotifier struct{}

func (t *TelegramNotifier) NotifyMessage(message string) {
	chatId := config.Get("CHAT_ID")

	id, err := strconv.Atoi(chatId)
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}

	api := notifier.GetApi()
	_, err = api.Send(tgbotapi.NewMessage(int64(id), message))
	if err != nil {
		fmt.Println("Error sending message", err)
	}
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
		log.Fatal(err)
	}
	err = writer.WriteField("chat_id", config.Get("CHAT_ID"))
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Infoln("Error sending photo:", err)
	} else {
		fmt.Println("Resp: ", res)
	}
}
