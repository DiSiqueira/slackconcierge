package main

import (
	"github.com/disiqueira/SlackConcierge/internal/cmd"
	"github.com/disiqueira/SlackConcierge/internal/infrastructure"
)

func main() {
	container := &infrastructure.Container{}
	command := cmd.RunCommand{
		LogService: container.LogService(),
		SlackService: container.SlackService(),
		HandlerService: container.HandlerService(),
	}

	command.Execute()
}
