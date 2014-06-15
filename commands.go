package main

import (
  "./commands"
  "github.com/codegangsta/cli"
)

func CommandsCollection() []cli.Command {
  return []cli.Command{
    commands.HookCommand,
    commands.LoadCommand,
  }
}
