//go:generate mockgen -source message.go -destination message_mock.go -package domain
package domain

type (
	Message struct {
		Text    string
		Channel string
	}

	MessageRepository interface {
		Consume() <-chan interface{}
		Post(message *Message)
	}
)
