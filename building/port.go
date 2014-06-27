package building

import (
	"fmt"
	"strings"

	"github.com/openit-lib/docker-exec/config"
)

type Port struct {
	command config.Command
}

func (p Port) IsPresent() bool {
	return len(p.command.Ports) > 0

}

func (p Port) Build() string {
	partial := []string{}
	for _, port := range p.command.Ports {
		partial = append(partial, fmt.Sprintf("-p %s", port))
	}
	return strings.Join(partial, " ")
}
