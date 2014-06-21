package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var HookCommand = cli.Command{
	Name:  "hook",
	Usage: "direnv hook <shell_name>",
	Action: func(c *cli.Context) {
		args := c.Args()
		current_shell, shell_err := CurrentShell(args)
		if shell_err != nil {
			fmt.Println(shell_err)
			return
		}

		fmt.Println((*current_shell).Hook())
	},
}
