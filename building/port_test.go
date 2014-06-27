package building

import (
	"testing"

	"github.com/openit-lib/docker-exec/config"
)

func TestBuildPortEmpty(t *testing.T) {
	cmd := config.Command{}
	port := Port{cmd}
	if port.IsPresent() {
		t.Error("Should not be present")
	}
}

func TestBuildPortSingle(t *testing.T) {
	port_values := []string{"8080"}
	expected_string := "-p 8080"
	cmd := config.Command{Ports: port_values}
	ports := Port{cmd}
	if !ports.IsPresent() {
		t.Error("Should be present")
	}

	if expected_string != ports.Build() {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_string, ports.Build())
	}
}

func TestBuildPortMultiple(t *testing.T) {
	port_values := []string{"8080", ":8081", "8082:8082"}
	expected_string := "-p 8080 -p :8081 -p 8082:8082"
	cmd := config.Command{Ports: port_values}
	ports := Port{cmd}
	if !ports.IsPresent() {
		t.Error("Should be present")
	}

	if expected_string != ports.Build() {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_string, ports.Build())
	}
}
