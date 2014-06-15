package env

import (
  "../config"
  "../shell"
  "testing"
)

func TestToShell(t *testing.T) {
  storage := Diff{make(map[string]string), make(map[string]string)}
  storage.Prev["a"] = "a_alias"
  storage.Prev["b"] = "b_alias"
  storage.Prev["c"] = "c_alias"

  storage.Next["c"] = "c_alias"
  storage.Next["d"] = "d_alias"
  storage.Next["e"] = "e_alias"

  shell_string := `unalias a;
unalias b;
alias c="c_alias";
alias d="d_alias";
alias e="e_alias";`
  storage_string := storage.ToShell(shell.BASH)

  if shell_string != storage_string {
    t.Errorf("Unexpected string\n-e: %s\n-o:%s", storage_string, shell_string)
  }
}

func TestLoadDiff(t *testing.T) {
  c1, err := config.LoadConfig("../test/envc1.yml")
  if err != nil {
    t.Error(err)
  }
  c2, err := config.LoadConfig("../test/envc2.yml")
  if err != nil {
    t.Error(err)
  }

  diff := LoadDiff(c1.Commands, c2.Commands)
  if diff == nil {
    t.Error("Can not load diff")
  }

  if len(diff.Prev) == 0 || len(diff.Next) == 0 {
    t.Error("Diff not fully loaded")
  }
}
