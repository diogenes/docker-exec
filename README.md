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
alias node=docker run --rm -it --volumes-from data_container -p 8080:8080 -v /vagrant:/vagrant ubuntu node
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

### Fish
Add the following line at the end of your "~/.config/fish/config.fish" file:

`eval (docker-exec hook fish)`


## Usage

In target directory create ".denv" file and add some command defenitions:

```yaml
commands:
  bundle: &base_cmd
    directory: "/vagrant"
    image: "devcenter"
    from: "bundles"
    link:
      - "postgresql:pg"
    dns:
      - 10.40.0.1
    volumes:
      - "/vagrant:/vagrant"
  rake: &bundled
    <<: *base_cmd
    prepend: "bundle exec"
  rspec:
    <<: *bundled
  pry:
    <<: *bundled
  rails: &rails
    <<: *bundled
    ports:
      - "8080:8080"
  unicorn:
    <<: *rails
```
Run `alias | grep docker` and you will see your aliases:
```bash
alias pry="docker run --rm -it -v /vagrant:/vagrant --dns 10.10.30.0 --link postgresql:pg -w /vagrant devcenter bundle exec pry";
alias bundle="docker run --rm -it -v /vagrant:/vagrant --dns 10.10.30.0 --link postgresql:pg -w /vagrant devcenter bundle";
alias rake="docker run --rm -it -v /vagrant:/vagrant --dns 10.10.30.0 --link postgresql:pg -w /vagrant devcenter bundle exec rake";
alias unicorn="docker run --rm -it -p 8080:8080 -v /vagrant:/vagrant --dns 10.10.30.0 --link postgresql:pg -w /vagrant devcenter bundle exec unicorn";
alias rspec="docker run --rm -it -v /vagrant:/vagrant --dns 10.10.30.0 --link postgresql:pg -w /vagrant devcenter bundle exec rspec";
alias rails="docker run --rm -it -p 8080:8080 -v /vagrant:/vagrant --dns 10.10.30.0 --link postgresql:pg -w /vagrant devcenter bundle exec rails";
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
