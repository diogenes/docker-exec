package env

import (
  "../config"
  "../shell"
  "fmt"
)

type Env struct {
  Commands map[string]*config.Command
  Shell    shell.Shell
}

func (e Env) Load() error {
  for k, v := range e.Commands {
    alias_string := e.Shell.Alias(k, v.Command())
    e.Run(alias_string)
  }
  return nil
}

func (e Env) Unload() error {
  for k, _ := range e.Commands {
    unalias_string := e.Shell.Unalias(k)
    e.Run(unalias_string)
  }
  return nil
}

func (e Env) Run(cmd string) error {
  fmt.Println(cmd)
  return nil
}
