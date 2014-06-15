package shell

import "testing"

func TestDetectShell(t *testing.T) {
  if DetechShell("unknown") != nil {
    t.Error("should be nil shell")
  }
}
