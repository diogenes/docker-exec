package config

import "testing"

func TestCommandCommand(t *testing.T) {
  expected_command := Command{"rails", "bundle exec", "/vagrant", "ruby", "-v /vagrant:/vagrant"}.Command()
  original_command := "docker run -it --rm -v /vagrant:/vagrant ruby cd /vagrant && bundle exec rails"
  if original_command != expected_command {
    t.Error("Command missmatch:", expected_command)
  }
}
