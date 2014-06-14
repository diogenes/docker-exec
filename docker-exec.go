package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Docker Exec"
	app.Author = "Alexander Kirillov"
	app.Run(os.Args)
}
