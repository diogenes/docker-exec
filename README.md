# docker-exec

[![Build Status](https://drone.io/github.com/openit-lib/docker-exec/status.png)](https://drone.io/github.com/openit-lib/docker-exec/latest)

Wrap, configure and execute commands in docker containers from console.

Docker - a great tool for developers and devops. But when you run the containers with many options, it becomes like hell.
Here is an example of running the usual container:
```
$ [host] > docker run --rm -it --volumes-from data_container -p 8080:8080 -v /vagrant:/vagrant -w /vagrant ubuntu /bin/bash
$ [docker] > any command in container
```
Repeat this several times with different parameters. Still not tired?

## Solution

Docker-exec allows you to automatically generate aliases for the necessary commands that you run in containers. For example:
```bash
export alias node=docker run --rm -it --volumes-from data_container -p 8080:8080 -v /vagrant:/vagrant ubuntu node
$ [host] > node install bower -g
```
This alias runs `node install bower` inside the docker container and cleanup container after execute.

## Install

### From binary package

1. Download binary package from [releases](https://github.com/openit-lib/docker-exec/releases)
2. Unpackage it
3. Run `make install` from package directory

or compile from source:

```bash
git clone https://github.com/openit-lib/docker-exec.git
cd docker-exec
make deps && make install
```
## Setup

### Bash

Add the following line at the end of your "~/.bashrc" file:

`eval "$(docker-exec hook bash)"`

### Zsh

Add the following line at the end of your "~/.zshrc" file:

`eval "$(docker-exec hook zsh)"`

## Usage

In target directory create ".denv" file and add some command defenitions:

```yaml
commands:
  npm:
    directory: "/vagrant"
    image: "node"
    args: "--rm -v /vagrant:/vagrant"
  rails: &rails_development
    prepend: "bundle exec"
    directory: "/vagrant"
    image: "ruby"
    args: "--rm -v /vagrant:/vagrant"
  rake:
    <<: *rails_development
  rspec:
    <<: *rails_development
```
Run `alias | grep docker` and you will see your aliases:
```bash
npm='docker run -it --rm --rm -v /vagrant:/vagrant -w /vagrant node npm'
rails='docker run -it --rm --rm -v /vagrant:/vagrant -w /vagrant ruby bundle exec rails'
rake='docker run -it --rm --rm -v /vagrant:/vagrant -w /vagrant ruby bundle exec rake'
rspec='docker run -it --rm --rm -v /vagrant:/vagrant -w /vagrant ruby bundle exec rspec'
```

Now you can run `npm --version`, `npm` runs inside new container and cleanup atfer run.

Go to another direcotry and repeat `alias | grep docker`.
You'll see that they're gone. That's right, there are only aliases in the directory in which a file `.denv` present.

## Contribute

Bug reports, contributions and forks are welcome.

All bugs or other forms of discussion happen on
<https://github.com/openit-lib/docker-exec/issues>

I express my special thanks to project [Direnv](https://github.com/zimbatm/direnv) for the idea, which I learned from their code.

## COPYRIGHT

Copyright (C) 2014 shared by all
[contributors](https://github.com/openit-lib/docker-exec/graphs/contributors) under
the MIT licence.
