package building

import (
	"testing"

	"github.com/openit-lib/docker-exec/config"
)

func TestBuildEnvEmpty(t *testing.T) {
	cmd := config.Command{}
	env := Environmet{cmd}
	if env.IsPresent() {
		t.Error("Should not be present")
	}
}

func TestBuildEnvSingle(t *testing.T) {
	env_values := map[string]string{
		"TEST": "value",
	}
	cmd := config.Command{Env: env_values}
	env := Environmet{cmd}
	if !env.IsPresent() {
		t.Error("Should be present")
	}
	if "-e TEST=value" != env.Build() {
		t.Error("Unexpected value")
	}
}
