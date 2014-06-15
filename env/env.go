package env

import (
  "../config"
  "../shell"
)

type Env struct {
  Commands map[string]*config.Command
  Shell    shell.Shell
}

func (e Env) Load() error {
  return nil
}

func (e Env) Uload() error {
  return nil
}
