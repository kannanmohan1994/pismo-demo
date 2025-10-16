package notifier

type NoopNotifier struct{}

func NewNoop() Notifier                                           { return NoopNotifier{} }
func (NoopNotifier) Notify(topic string, data []byte) (err error) { return }
