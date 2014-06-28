package building

import (
	"testing"

	"github.com/openit-lib/docker-exec/config"
)

func TestBuildPartial(t *testing.T) {
	command := config.Command{
		Name:      "bash",
		Directory: "/vagrant",
		Image:     "ubuntu",
		Volumes:   []string{"/vagrant:/vagrant"},
	}
	expected_command := "docker run --rm -it -v /vagrant:/vagrant -w /vagrant ubuntu bash"
	real_command := NewAliasBuilder(&command).Build()
	if expected_command != real_command {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_command, real_command)
	}
}

func TestBuildFull(t *testing.T) {
	command := config.Command{
		Name:      "unicorn",
		Prepend:   "bundle exec",
		Directory: "/vagrant",
		Image:     "ruby",
		Link:      "postgresql",
		Args:      "-E production",
		Ports:     []string{"8080", "3000"},
		Volumes:   []string{"/vagrant:/vagrant"},
		From:      []string{"bundles"},
		Dns:       []string{"8.8.8.8"},
		Env:       map[string]string{"USER": "Kirillov"},
	}
	expected_command := "docker run --rm -it -p 8080 -p 3000 -v /vagrant:/vagrant --from-image bundles --dns 8.8.8.8 --link postgresql -w /vagrant -e USER=Kirillov ruby bundle exec unicorn -E production"
	real_command := NewAliasBuilder(&command).Build()
	if expected_command != real_command {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_command, real_command)
	}
}
