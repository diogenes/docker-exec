package env

import (
  "github.com/openit-lib/docker-exec/config"
  "github.com/openit-lib/docker-exec/shell"
  "testing"
)

func TestToShell(t *testing.T) {
  c1, err := config.LoadConfig("../test/envc1.yml")
  if err != nil {
    t.Error(err)
  }

  diff := Diff{make(map[string]string), make(map[string]string), c1}
  diff.Prev["a"] = "a_alias"
  diff.Prev["b"] = "b_alias"
  diff.Prev["c"] = "c_alias"

  diff.Next["c"] = "c_alias"
  diff.Next["d"] = "d_alias"
  diff.Next["e"] = "e_alias"

  shell_string := `unalias a;
unalias b;
alias c="c_alias";
alias d="d_alias";
alias e="e_alias";
export _DENV=eJxcjcsKwjAQRX8lzLrBfXaiGzfiL0ybYQjkIdMHlpJ_d0pRopuBc7lzzwaXkhJmP4LbgMt-75gI3A4dPISelL2iwjUIDVORVfHExQfR8JaQj3rEzBqchXUMrJUEtQPBEMdm9uB2uZ-zj2ToRcO_ZEEWzFOjkblffyXGLuZTdN-PWus7AAD__9QPRik=`
  storage_string := diff.ToShell(shell.BASH)

  if shell_string != storage_string {
    t.Errorf("Unexpected string\n-e:%s\n-o:%s", storage_string, shell_string)
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

  diff := LoadDiff(c1, c2)
  if diff == nil {
    t.Error("Can not load diff")
  }

  if len(diff.Prev) == 0 || len(diff.Next) == 0 {
    t.Error("Diff not fully loaded")
  }
}
