package config

import (
  "io/ioutil"

  "gopkg.in/yaml.v1"
)

type Config struct {
  path     string
  Commands map[string]*Command
}

func LoadConfig(path string) (config *Config, err error) {
  config = &Config{
    path: path,
  }

  config_buff, err := ioutil.ReadFile(path)

  if err != nil {
    return nil, err
  }

  if err = yaml.Unmarshal(config_buff, &config); err != nil {
    return nil, err
  }

  for k, v := range config.Commands {
    v.Name = k
  }

  return config, err
}
