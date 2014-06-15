package shell

type zsh int

var ZSH zsh

const ZSH_HOOK = `
_docker_exec_hook() {
  eval "$(docker-exec export zsh)";
}
typeset -a precmd_functions
if [[ -z $precmd_functions[(r)_docker_exec_hook] ]]; then
  precmd_functions+=_docker_exec_hook;
fi
`

func (z zsh) Hook() string {
  return ZSH_HOOK
}

func (z zsh) Alias(key, value string) string {
  return "alias " + key + "=\"" + value + "\";"
}

func (z zsh) Unalias(key string) string {
  return "unalias " + key + ";"
}
