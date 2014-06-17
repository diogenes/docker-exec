package env

import (
  "../config"
  "os"
  "testing"
)

func TestSessionConfig(t *testing.T) {
  original_config, _ := config.LoadConfig("../test/envc1.yml")
  dumped_string := StoreSession(original_config)
  os.Setenv("_DENV", dumped_string)
  expected_config, _ := SessionConfig()

  for key, v := range original_config.Commands {
    if expected_config.Commands[key].Name != v.Name {
      t.Error("expected config not eql to original config")
    }
  }

}

func TestStoreSession(t *testing.T) {
  c1, err := config.LoadConfig("../test/envc1.yml")
  if err != nil {
    t.Error(err)
  }
  StoreSession(c1)
}
