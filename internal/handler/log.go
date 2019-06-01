package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/disiqueira/SlackConcierge/internal/domain"
)

type (
	Log struct {
	}
)

func (h Log) Execute(message *domain.Message) ([]*domain.Message, error) {
	log.Println(fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), message.Text))
	return nil, nil
}
