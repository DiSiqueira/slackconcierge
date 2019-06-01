//go:generate mockgen -source handler.go -destination handler_mock.go -package handler
package handler

import "github.com/disiqueira/SlackConcierge/internal/domain"

type (
	Handler interface {
		Execute(message *domain.Message) ([]*domain.Message, error)
	}
)
