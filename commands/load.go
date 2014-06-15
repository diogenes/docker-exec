package commands

import (
  "fmt"
  "os"
  "path"

  "../config"
  "github.com/codegangsta/cli"
)

var LoadCommand = cli.Command{
  Name:  "load",
  Usage: "direnv load",
  Action: func(c *cli.Context) {
    args := c.Args()
    var current_path string

    if len(args) > 0 {
      current_path = args[0]
    } else {
      current_path, _ = os.Getwd()
    }

    current_path = ExpandPath(current_path)

    app_config, err := config.LoadConfig(path.Join(current_path, ".denv"))
    if err != nil {
      fmt.Printf("Error: %s", err)
      return
    }

    fmt.Printf("config file = %v", app_config)
    return
  },
}

func ExpandPath(conf_path string) string {
  if path.IsAbs(conf_path) {
    return conf_path
  }

  relative_to, _ := os.Getwd()
  return path.Join(relative_to, conf_path)
}
