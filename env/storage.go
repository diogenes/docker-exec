package env

import (
  "../config"
  "os"
)

func SessionConfig() *config.Config {
  dumped := os.Getenv("_DENV")
  if dumped == "" {
    return nil
  }
  return LoadBase64String(dumped)
}

func LoadBase64String(serialized_string string) *config.Config {
  config := config.Config{}
  unmarshal(serialized_string, &config)
  return &config
}

func StoreSession(c *config.Config) string {
  base_64_string := marshal(c)
  return base_64_string
}