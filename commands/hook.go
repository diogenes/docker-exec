package commands

import (
  "fmt"

  "../shell"
  "github.com/codegangsta/cli"
)

var HookCommand = cli.Command{
  Name:  "hook",
  Usage: "direnv hook <shell_name>",
  Action: func(c *cli.Context) {
    args := c.Args()
    if len(args) == 0 {
      fmt.Println("Unknown shell")
      return
    }
    target := args[0]

    current_shell := shell.DetechShell(target)
    if current_shell == nil {
      fmt.Errorf("Unknown shell <%s>", target)
    }

    fmt.Println(current_shell.Hook())
  },
}
