package building

import "strings"

type List struct {
	prefix  string
	options []string
}

func NewListBuilder(prefix string, options []string) List {
	return List{prefix, options}
}

func (l List) IsPresent() bool {
	return len(l.options) > 0
}

func (l List) Build() string {
	partial := []string{}
	for _, opt := range l.options {
		partial = append(partial, NewSingleBuilder(l.prefix, opt).Build())
	}
	return strings.Join(partial, " ")
}
