package env

import (
  "../config"
  "../shell"
)

type Env struct {
  Commands map[string]*config.Command
  Shell    shell.Shell
}
