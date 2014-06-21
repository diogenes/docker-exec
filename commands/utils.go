package commands

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/openit-lib/docker-exec/config"
	"github.com/openit-lib/docker-exec/shell"

	"github.com/codegangsta/cli"
)

func CurrentConfig(args cli.Args) (*config.Config, error) {
	if len(args) == 0 {
		return nil, errors.New("Config not specify")
	}

	current_path, _ := os.Getwd()
	current_path = ExpandPath(current_path)

	app_config, err := config.LoadConfig(path.Join(current_path, ".denv"))
	if err != nil {
		return nil, err
	}
	return app_config, nil
}

func CurrentShell(args cli.Args) (*shell.Shell, error) {
	if len(args) == 0 {
		return nil, errors.New("Shell not specify")
	}
	current_shell := shell.DetectShell(args[0])
	if current_shell == nil {
		return nil, fmt.Errorf("Shell '%s' not found", args[0])
	}
	return &current_shell, nil
}

func ExpandPath(conf_path string) string {
	if path.IsAbs(conf_path) {
		return conf_path
	}

	relative_to, _ := os.Getwd()
	return path.Join(relative_to, conf_path)
}
