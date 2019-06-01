//go:generate mockgen -source handler.go -destination handler_mock.go -package application
package application

import (
	"github.com/disiqueira/SlackConcierge/internal/domain"
	"github.com/disiqueira/SlackConcierge/internal/handler"
)

type (
	handlerService struct {
		handlerList []handler.Handler
	}
	HandlerService interface {
		Handle(msg *domain.Message) ([]*domain.Message, []error)
	}
)

func NewHandlerService(handlerList []handler.Handler) HandlerService {
	return &handlerService{
		handlerList: handlerList,
	}
}

func (s *handlerService) Handle(msg *domain.Message) ([]*domain.Message, []error) {
	var responses []*domain.Message
	var errors []error

	for _, h := range s.handlerList {
		response, err := h.Execute(msg)
		responses = append(responses, response...)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return responses, errors
}
