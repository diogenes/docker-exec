package building

import (
	"testing"

	"github.com/openit-lib/docker-exec/config"
)

func TestBuildListEmpty(t *testing.T) {
	cmd := config.Command{}
	list := NewListBuilder("-p", cmd.Ports)
	if list.IsPresent() {
		t.Error("should be empty")
	}
}

func TestBuildListDingle(t *testing.T) {
	expected_string := "-p 8080"
	list_values := []string{"8080"}
	cmd := config.Command{Ports: list_values}
	list := NewListBuilder("-p", cmd.Ports)
	if !list.IsPresent() {
		t.Error("Should be present")
	}

	real_string := list.Build()
	if expected_string != real_string {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_string, real_string)
	}
}
