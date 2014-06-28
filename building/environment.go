package building

import (
	"fmt"
	"strings"

	"github.com/openit-lib/docker-exec/config"
)

type Environmet struct {
	command config.Command
}

func NewEnvironmantBuilder(command config.Command) Environmet {
	return Environmet{command}
}

func (e Environmet) IsPresent() bool {
	return len(e.command.Env) > 0

}

func (e Environmet) Build() string {
	partial := []string{}
	for key, value := range e.command.Env {
		partial = append(partial, fmt.Sprintf("-e %s=%s", key, value))
	}
	return strings.Join(partial, " ")
}
