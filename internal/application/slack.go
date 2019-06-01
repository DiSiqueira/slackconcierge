//go:generate mockgen -source slack.go -destination slack_mock.go -package application
package application

import (
	"github.com/disiqueira/SlackConcierge/internal/domain"
)

type (
	slackService struct {
		messageRepository domain.MessageRepository
		bufferSize        int
	}

	SlackService interface {
		Messages() chan *domain.Message
		Post(message *domain.Message)
		PostBulk(msgList []*domain.Message)
	}
)

func NewSlackService(repository domain.MessageRepository, bufferSize int) SlackService {
	return &slackService{
		messageRepository: repository,
		bufferSize:        bufferSize,
	}
}

func (s *slackService) Messages() chan *domain.Message {
	msgList := make(chan *domain.Message, s.bufferSize)
	go func() {
		for ev := range s.messageRepository.Consume() {
			switch msg := ev.(type) {
			case *domain.Message:
				msgList <- msg
			}
		}
	}()
	return msgList
}

func (s *slackService) Post(message *domain.Message) {
	s.messageRepository.Post(message)
}

func (s *slackService) PostBulk(msgList []*domain.Message) {
	if len(msgList) == 0 {
		return
	}
	s.Post(msgList[0])
	s.PostBulk(msgList[1:])
}
