package notifier

type Notifier interface {
	Notify(topic string, data []byte) error
}
