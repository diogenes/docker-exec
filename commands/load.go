package commands

import (
  "fmt"

  . "../env"
  "github.com/codegangsta/cli"
)

var LoadCommand = cli.Command{
  Name:  "load",
  Usage: "direnv load <shell name>",
  Action: func(c *cli.Context) {
    args := c.Args()
    app_config, config_err := CurrentConfig(args)
    current_shell, shell_err := CurrentShell(args)

    if shell_err != nil {
      fmt.Println(shell_err)
      return
    }
    if config_err != nil {
      fmt.Println(config_err)
      return
    }

    Env{app_config.Commands, *current_shell}.Load()
    return
  },
}
