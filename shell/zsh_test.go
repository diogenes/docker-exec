package shell

import "testing"

func TestDetechZsh(t *testing.T) {
	shell := DetectShell("zsh")
	if shell == nil {
		t.Errorf("Shell <%s> not found", shell)
	}
}

func TestZshAlias(t *testing.T) {
	shell := DetectShell("zsh")
	original_command := "alias rails=\"bundle exec rails\";"
	expected_command := shell.Alias("rails", "bundle exec rails")
	if original_command != expected_command {
		t.Errorf("error alias build '%s'", expected_command)
	}
}

func TestZshUnalias(t *testing.T) {
	shell := DetectShell("zsh")
	original_command := "unalias rails;"
	expected_command := shell.Unalias("rails")
	if original_command != expected_command {
		t.Errorf("error alias build '%s'", expected_command)
	}
}
