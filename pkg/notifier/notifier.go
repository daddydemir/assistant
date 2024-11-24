package notifier

import "github.com/daddydemir/assistant/pkg/notifier/telegram"

type Notifier interface {
	NotifyMessage(message string)
	NotifyImage(imagePath string)
}

func GetActiveNotifier() Notifier {
	return &telegram.TelegramNotifier{}
}
