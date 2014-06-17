package commands

import (
  "fmt"

  "../config"
  "../env"
  "github.com/codegangsta/cli"
)

var RcCommand = cli.Command{
  Name:  "rc",
  Usage: "direnv rc",
  Action: func(c *cli.Context) {
    args := c.Args()

    next_config, err := config.LoadConfig("./.denv")
    if err != nil {
      next_config = &config.Config{}
    }
    prev_config, err := env.SessionConfig()
    if err != nil {
      prev_config = &config.Config{}
    }

    diff := env.LoadDiff(prev_config.Commands, next_config.Commands)
    if diff == nil {
      return
    }

    current_shell, shell_err := CurrentShell(args)
    if shell_err != nil {
      fmt.Println(shell_err)
      return
    }

    fmt.Println(diff.ToShell(*current_shell))
    return
  },
}
