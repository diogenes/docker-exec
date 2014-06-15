package shell

import (
  "path/filepath"
)

type Shell interface {
  Hook() string
  Alias(key, value string) string
  Unalias(key string) string
}

func DetechShell(name string) Shell {
  name = filepath.Base(name)
  // $0 starts with "-"
  if name[0:1] == "-" {
    name = name[1:]
  }
  switch name {
  case "bash":
    return BASH
  case "zsh":
    return ZSH
  case "fish":
    return nil
  }
  return nil
}
