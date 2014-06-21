package shell

import "testing"

func TestDetechBash(t *testing.T) {
	shell := DetectShell("bash")
	if shell == nil {
		t.Errorf("Shell <%s> not found", shell)
	}
}

func TestBashAlias(t *testing.T) {
	shell := DetectShell("bash")
	original_command := "alias rails=\"bundle exec rails\";"
	expected_command := shell.Alias("rails", "bundle exec rails")
	if original_command != expected_command {
		t.Errorf("error alias build '%s'", expected_command)
	}
}

func TestBashUnalias(t *testing.T) {
	shell := DetectShell("bash")
	original_command := "unalias rails;"
	expected_command := shell.Unalias("rails")
	if original_command != expected_command {
		t.Errorf("error alias build '%s'", expected_command)
	}
}
