package telegram

import (
	"bytes"
	"fmt"
	"github.com/daddydemir/assistant/internal/log"
	"github.com/daddydemir/assistant/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

func SendMessage(message string) {
	chatId := config.Get("CHAT_ID")

	idInt, _ := strconv.Atoi(chatId)
	_, err := config.Bot.Send(tgbotapi.NewMessage(int64(idInt), message))
	if err != nil {
		log.Errorln(err)
	}

}

func SendImage(filename string) {
	path := config.Get("IMAGE_FOLDER") + filename
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
	_, err = client.Do(req)
	if err != nil {
		log.Infoln("Error sending photo:", err)
	}
}
