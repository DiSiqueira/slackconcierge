package thirdparty

import (
	"github.com/disiqueira/SlackConcierge/internal/domain"

	"github.com/nlopes/slack"
)

type (
	SlackRepository struct {
		client     *slack.Client
		rtm        *slack.RTM
		bufferSize int
	}
)

func NewSlackRepository(token string, bufferSize int) (*SlackRepository, error) {
	s := &SlackRepository{
		bufferSize: bufferSize,
	}
	s.client = slack.New(token)
	s.rtm = s.client.NewRTM()
	go s.rtm.ManageConnection()
	return s, nil
}

func (s *SlackRepository) Consume() <-chan interface{} {
	eventList := make(chan interface{}, s.bufferSize)
	go func() {
		for msg := range s.rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				eventList <- &domain.Message{
					Text: ev.Text,
					Channel: ev.Channel,
				}
			}
		}
	}()
	return eventList
}

func (s *SlackRepository) Post(message *domain.Message) {
	s.rtm.SendMessage(s.rtm.NewOutgoingMessage(message.Text, message.Channel))
}
