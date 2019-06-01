package handler

import (
	"fmt"
	"github.com/disiqueira/SlackConcierge/internal/domain"
	"github.com/disiqueira/SlackConcierge/internal/program"
	"github.com/urfave/cli"
	"strings"
)

type (
	Program struct {
	}
)

func (h Program) Execute(message *domain.Message) ([]*domain.Message, error) {
	if message.Text[0] != '#' {
		fmt.Println("message ignored")
		return nil, nil
	}

	cmd := cli.NewApp()
	cmd.Action = func(c *cli.Context) error {
		return cli.ShowAppHelp(c)
	}
	cmd.Commands = program.Commands()

	fmt.Println(strings.Fields(string(message.Text[1:])))
	if err := cmd.Run(strings.Fields(string(message.Text[1:]))); err != nil {
		return nil, err
	}

	return nil, nil
}
