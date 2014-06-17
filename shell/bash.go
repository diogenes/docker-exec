package shell

import (
  "../config"
)

type bash int

var BASH bash

const BASH_HOOK = `
_docker_exec_hook() {
  eval "$(docker-exec rc bash)";
};
if ! [[ "$PROMPT_COMMAND" =~ _docker_exec_hook ]]; then
  PROMPT_COMMAND="_docker_exec_hook;$PROMPT_COMMAND";
fi
`

func (b bash) Hook() string {
  return BASH_HOOK
}

func (b bash) Denv(current_config *config.Config) string {
  return ""
}

func (b bash) Alias(key, value string) string {
  return "alias " + key + "=\"" + value + "\";"
}

func (b bash) Unalias(key string) string {
  return "unalias " + key + ";"
}
