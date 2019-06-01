package program

import (
	"github.com/disiqueira/SlackConcierge/internal/program/hello/cmd"
	"github.com/urfave/cli"
)

func Commands() []cli.Command {
	return []cli.Command{
		cmd.New(),
	}
}
