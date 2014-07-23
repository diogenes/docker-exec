package env

import (
	"testing"

	"github.com/openit-lib/docker-exec/building"
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

func TestBuldCommand(t *testing.T) {
	expected_rails := "docker run --rm -it -v /vagrant -w /vagrant ruby bundle exec rails"
	expected_cucumber := "docker run --rm -it -p 8080 -p 8081 -v /vagrant --link postgresql:db --volumes-from bundles -w /vagrant -e USER=Kirillov -e RAILS_ENV=test ruby bundle exec cucumber"
	cnf, err := config.LoadConfig("../test/denv.yml")
	if err != nil {
		t.Error(err)
	}
	real_rails := building.NewAliasBuilder(cnf.Commands["rails"]).Build()
	if expected_rails != real_rails {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_rails, real_rails)
	}

	real_cucumber := building.NewAliasBuilder(cnf.Commands["cucumber"]).Build()
	if expected_cucumber != real_cucumber {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_cucumber, real_cucumber)
	}
}

func TestBuildWithAlias(t *testing.T) {
	expected_cmd := "docker run --rm -it ubuntu /bin/bash"

	cnf, err := config.LoadConfig("../test/denv.yml")
	if err != nil {
		t.Error(err)
	}
	real_cmd := building.NewAliasBuilder(cnf.Commands["do_bash"]).Build()
	if expected_cmd != real_cmd {
		t.Errorf("Unexpected value\n-o:%s\n-e:%s", expected_cmd, real_cmd)
	}
}
