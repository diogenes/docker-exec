package shell

import "testing"

func TestDetectShell(t *testing.T) {
	if DetectShell("unknown") != nil {
		t.Error("should be nil shell")
	}
}
