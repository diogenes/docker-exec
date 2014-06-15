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
    fmt.Println(alias_string)
    //e.Run(alias_string)
  }
  return nil
}

func (e Env) Uload() error {
  for k, _ := range e.Commands {
    unalias_string := e.Shell.Unalias(k)
    fmt.Println(unalias_string)
  }
  return nil
}
