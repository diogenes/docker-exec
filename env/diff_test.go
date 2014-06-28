package env

import (
	"testing"

	"github.com/openit-lib/docker-exec/config"
)

func TestLoadDiff(t *testing.T) {
	c1, err := config.LoadConfig("../test/envc1.yml")
	if err != nil {
		t.Error(err)
	}

	c2, err := config.LoadConfig("../test/envc2.yml")
	if err != nil {
		t.Error(err)
	}

	diff := LoadDiff(c1, c2)
	if diff == nil {
		t.Error("Can not load diff")
	}

	if len(diff.Prev) == 0 || len(diff.Next) == 0 {
		t.Error("Diff not fully loaded")
	}
}
