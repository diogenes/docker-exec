package config

type Command struct {
	Name      string
	Cmd       string
	Prepend   string
	Directory string
	Image     string
	From      string
	Args      string
	Link      []string
	Ports     []string
	Volumes   []string
	Dns       []string
	Env       map[string]string
}
