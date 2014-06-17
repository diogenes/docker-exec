package env

import (
  "../config"
  "../shell"
  "strings"
)

type Diff struct {
  Prev          map[string]string
  Next          map[string]string
  CurrentConfig *config.Config
}

func LoadDiff(prev_conf, next_conf *config.Config) *Diff {
  storage := &Diff{make(map[string]string), make(map[string]string), next_conf}

  c1 := prev_conf.Commands
  c2 := next_conf.Commands

  in := func(key string, e map[string]*config.Command) bool {
    _, ok := e[key]
    return ok
  }

  for key := range c1 {
    if c2[key] != c1[key] || !in(key, c2) {
      storage.Prev[key] = c1[key].Command()
    }
  }

  for key := range c2 {
    if c2[key] != c1[key] || !in(key, c1) {
      storage.Next[key] = c2[key].Command()
    }
  }

  return storage
}

func (self Diff) ToShell(sh shell.Shell) string {
  str := []string{}
  for key := range self.Prev {
    _, ok := self.Next[key]
    if !ok {
      str = append(str, sh.Unalias(key))
    }
  }

  for key, value := range self.Next {
    str = append(str, sh.Alias(key, value))
  }

  if self.CurrentConfig != nil {
    str = append(str, sh.Denv(self.CurrentConfig))
  }

  return strings.Join(str, "\n")
}
