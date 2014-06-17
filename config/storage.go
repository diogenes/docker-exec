package config

import (
  "errors"
  "os"
)

func SessionConfig() (*Config, error) {
  dumped := os.Getenv("_DENV")
  if dumped == "" {
    return nil, errors.New("empty session")
  }
  return LoadBase64String(dumped), nil
}

func LoadBase64String(serialized_string string) *Config {
  config := Config{}
  unmarshal(serialized_string, &config)
  return &config
}

func StoreSession(c *Config) string {
  base_64_string := marshal(c)
  return base_64_string
}
