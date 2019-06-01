package cmd

import (
	"github.com/disiqueira/SlackConcierge/internal/application"
)

type (
	RunCommand struct {
		SlackService   application.SlackService
		HandlerService application.HandlerService
		LogService     application.LogService
	}
)

func (r *RunCommand) Execute() {
	for msg := range r.SlackService.Messages() {
		responses, errors := r.HandlerService.Handle(msg)
		if len(errors) > 0 {
			r.LogService.Errors(errors)
		}
		if len(responses) > 0 {
			r.SlackService.PostBulk(responses)
		}
	}
}
