package building

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/openit-lib/docker-exec/config"
)

type PartBuilder interface {
	Build() string
	IsPresent() bool
}

type AliasBuilder struct {
	cmd, name, prepend, image, args string
	builders                        []PartBuilder
}

func NewAliasBuilder(command *config.Command) *AliasBuilder {
	builders := []PartBuilder{
		NewListBuilder("-p", command.Ports),
		NewListBuilder("-v", command.Volumes),
		NewListBuilder("--dns", command.Dns),
		NewListBuilder("--link", command.Link),
		NewSingleBuilder("--volumes-from", command.From),
		NewSingleBuilder("-w", command.Directory),
		NewEnvironmantBuilder(*command),
	}

	return &AliasBuilder{
		cmd:      command.Cmd,
		name:     command.Name,
		prepend:  command.Prepend,
		image:    command.Image,
		args:     command.Args,
		builders: builders,
	}
}

func (self AliasBuilder) Build() string {
	options := self.buildOptions()
	cmd := self.buildCmd()

	return fmt.Sprintf("docker run %s %s %s", options, self.image, cmd)
}

func (self AliasBuilder) buildOptions() string {
	options := []string{"--rm -it"}
	for _, builder := range self.builders {
		if builder.IsPresent() {
			options = append(options, builder.Build())
		}
	}
	return strings.Join(options, " ")
}

func (self AliasBuilder) buildCmd() string {
	cmd_line := bytes.NewBufferString("")
	if self.prepend != "" {
		cmd_line.WriteString(fmt.Sprintf("%s ", self.prepend))
	}
	if self.cmd != "" {
		cmd_line.WriteString(self.cmd)
	} else {
		cmd_line.WriteString(self.name)
	}
	if self.args != "" {
		cmd_line.WriteString(fmt.Sprintf(" %s", self.args))
	}
	return cmd_line.String()
}
