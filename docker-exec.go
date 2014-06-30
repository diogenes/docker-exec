package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Docker Exec"
	app.Author = "Alexander Kirillov"
	app.Commands = CommandsCollection()
	app.Version = "0.2.0"
	app.Run(os.Args)
}
