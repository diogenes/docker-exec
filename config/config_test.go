package config

import "testing"

func TestLoadConfig(t *testing.T) {
  config, err := LoadConfig("../test/denv.yml")
  if err != nil {
    t.Error(err)
  }

  if config.Commands == nil {
    t.Error("commands empty")
  }

  for k, v := range config.Commands {
    if k != v.Name {
      t.Error("Not eql key and command name")
    }
  }
}
