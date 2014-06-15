package config

import "testing"

func TestCommandCommand(t *testing.T) {
  expected_command := Command{"rails", "bundle exec", "/vagrant", "ruby", "-v /vagrant:/vagrant"}.Command()
  original_command := "docker run -it --rm -v /vagrant:/vagrant -w /vagrant ruby bundle exec rails"
  if original_command != expected_command {
    t.Errorf("Command missmatch:\n --expected: %s\n --original: %s", expected_command, original_command)
  }
}
