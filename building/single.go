package building

import "fmt"

type Single struct {
	prefix string
	option string
}

func (s Single) IsPresent() bool {
	return s.option != ""
}

func (s Single) Build() string {
	return fmt.Sprintf("%s %s", s.prefix, s.option)
}

func NewSingleBuilder(prefix, value string) Single {
	return Single{prefix, value}
}
