default: build

build: vet
	go build -v -o ./bin/docker-exec .

tests:
	go test ./...

deps:
	@( \
		go get github.com/codegangsta/cli; \
		go get launchpad.net/goyaml; \
	)

vet:
	go vet ./...
