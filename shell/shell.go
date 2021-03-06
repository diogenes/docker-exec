package shell

import (
	"path/filepath"

	"github.com/openit-lib/docker-exec/config"
)

type Shell interface {
	Hook() string
	Denv(current_config *config.Config) string
	Alias(key, value string) string
	Unalias(key string) string
}

func DetectShell(name string) Shell {
	name = filepath.Base(name)
	// $0 starts with "-"
	if name[0:1] == "-" {
		name = name[1:]
	}
	switch name {
	case "bash":
		return BASH
	case "zsh":
		return ZSH
	case "fish":
		return FISH
	}
	return nil
}
