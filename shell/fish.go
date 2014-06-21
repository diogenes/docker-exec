package shell

import (
	"github.com/openit-lib/docker-exec/config"
)

type fish int

var FISH fish

const FISH_HOOK = `
function __docker_exec_hook --on-event fish_prompt;
	eval (docker-exec rc fish);
end
`

func (f fish) Hook() string {
	return FISH_HOOK
}

func (f fish) Denv(current_config *config.Config) string {
	return "set -x _DENV " + current_config.StoreSession()
}

func (f fish) Alias(key, value string) string {
	return "alias " + key + "=\"" + value + "\";"
}

func (f fish) Unalias(key string) string {
	return "functions -e " + key + ";"
}
