package shell

import "testing"

func TestDetechFish(t *testing.T) {
	shell := DetectShell("fish")
	if shell == nil {
		t.Errorf("Shell <%s> not found", shell)
	}
}

func TestFishAlias(t *testing.T) {
	shell := DetectShell("zsh")
	original_command := "alias rails=\"bundle exec rails\";"
	expected_command := shell.Alias("rails", "bundle exec rails")
	if original_command != expected_command {
		t.Errorf("error alias build '%s'", expected_command)
	}
}

func TestFishUnalias(t *testing.T) {
	shell := DetectShell("fish")
	original_command := "functions -e rails;"
	expected_command := shell.Unalias("rails")
	if original_command != expected_command {
		t.Errorf("error alias build '%s'", expected_command)
	}
}
