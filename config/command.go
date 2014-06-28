package config

type Command struct {
	Alias     string
	Name      string
	Prepend   string
	Directory string
	Image     string
	Link      string
	Args      string
	Ports     []string
	Volumes   []string
	From      []string
	Dns       []string
	Env       map[string]string
}
