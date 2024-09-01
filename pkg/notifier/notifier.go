package notifier

type Notifier interface {
	NotifyMessage(message string)
	NotifyImage(imagePath string)
}
