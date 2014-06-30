package config

type Command struct {
	Name      string
	Cmd       string
	Prepend   string
	Directory string
	Image     string
	Link      string
	From      string
	Args      string
	Ports     []string
	Volumes   []string
	Dns       []string
	Env       map[string]string
}
