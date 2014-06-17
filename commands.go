package main

import (
  "github.com/codegangsta/cli"
  "github.com/openit-lib/docker-exec/commands"
)

func CommandsCollection() []cli.Command {
  return []cli.Command{
    commands.HookCommand,
    commands.RcCommand,
  }
}
