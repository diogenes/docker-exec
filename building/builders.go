package building

import "github.com/openit-lib/docker-exec/config"

type PartBuilder interface {
	Build() string
	IsPresent() bool
}

type AliasBuilder struct {
	builders []PartBuilder
}

func NewAliasBuilder(command *config.Command) *AliasBuilder {
	return nil
}
