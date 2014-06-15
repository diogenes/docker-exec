package commands

import (
  "fmt"

  . "../env"
  "github.com/codegangsta/cli"
)

var UnloadCommand = cli.Command{
  Name:  "unload",
  Usage: "direnv unload <shell_name>",
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

    Env{app_config.Commands, *current_shell}.Unload()
    return
  },
}
