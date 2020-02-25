package action

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/urfave/cli/v2"
)

// Echo ...
var Echo = cli.Command{
	Name:  "echo",
	Usage: "3... 2... 1... liftoff",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "d",
		},
	},
	Action: echo,
}

func echo(c *cli.Context) error {
	return nil
}
