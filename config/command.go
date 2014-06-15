package config

type Command struct {
  Name      string
  Prepend   string
  Directory string
  Image     string
  Args      string
}

func (c Command) Command() string {
  cmd := c.Name
  if c.Prepend != "" {
    cmd = c.Prepend + " " + cmd
  }

  docker_cli := "docker run -it --rm"
  if c.Args != "" {
    docker_cli = docker_cli + " " + c.Args
  }

  if c.Directory != "" {
    docker_cli = docker_cli + " -w " + c.Directory
  }

  if c.Image != "" {
    docker_cli = docker_cli + " " + c.Image
  }

  return docker_cli + " " + cmd
}
