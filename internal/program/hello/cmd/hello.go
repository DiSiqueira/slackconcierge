package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func New() cli.Command {
	return cli.Command{
		Name:  "hello",
		Usage: "Say Hello!",
		Aliases: []string{"h"},
		Action: func(c *cli.Context) error {
			fmt.Println("Helloaaa!")
			return nil
		},
	}
}
