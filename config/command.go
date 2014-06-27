package config

import (
	"bytes"
	"fmt"
)

type Command struct {
	ALias     string
	Name      string
	Prepend   string
	Directory string
	Image     string
	Args      string
	Ports     []string
	Volumes   []string
	From      []string
	Dns       []string
	Env       map[string]string
}

func (c Command) Command() string {
	cmd := bytes.NewBufferString("")
	if c.Prepend != "" {
		cmd.WriteString(fmt.Sprintf("%s ", c.Prepend))
	}
	cmd.WriteString(fmt.Sprintf("%s", c.Name))

	docker_cli := bytes.NewBufferString("docker run -it --rm")
	if c.Args != "" {
		docker_cli.WriteString(fmt.Sprintf(" %s", c.Args))
	}

	if c.Directory != "" {
		docker_cli.WriteString(fmt.Sprintf(" -w %s", c.Directory))
	}

	docker_cli.WriteString(fmt.Sprintf(" %s", c.Image))

	return fmt.Sprintf("%s %s", docker_cli, cmd)
}
