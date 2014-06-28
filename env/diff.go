package env

import (
	"strings"

	"github.com/openit-lib/docker-exec/building"
	"github.com/openit-lib/docker-exec/config"
	"github.com/openit-lib/docker-exec/shell"
)

type Diff struct {
	Prev          map[string]string
	Next          map[string]string
	CurrentConfig *config.Config
}

func LoadDiff(prev_conf, next_conf *config.Config) *Diff {
	storage := &Diff{make(map[string]string), make(map[string]string), next_conf}

	c1 := prev_conf.Commands
	c2 := next_conf.Commands

	in := func(key string, e map[string]*config.Command) bool {
		_, ok := e[key]
		return ok
	}

	for key := range c1 {
		if c2[key] != c1[key] || !in(key, c2) {
			storage.Prev[key] = building.NewAliasBuilder(c1[key]).Build()
		}
	}

	for key := range c2 {
		if c2[key] != c1[key] || !in(key, c1) {
			storage.Next[key] = building.NewAliasBuilder(c2[key]).Build()
		}
	}

	return storage
}

func (self Diff) ToShell(sh shell.Shell) string {
	str := []string{}
	for key := range self.Prev {
		_, ok := self.Next[key]
		if !ok {
			str = append(str, sh.Unalias(key))
		}
	}

	for key, value := range self.Next {
		str = append(str, sh.Alias(key, value))
	}

	if self.CurrentConfig != nil {
		str = append(str, sh.Denv(self.CurrentConfig))
	}

	return strings.Join(str, "\n")
}
