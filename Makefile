default: build

build: vet
	./build.sh

tests: vet
	go test ./...

deps:
	go get code.google.com/p/go.tools/cmd/vet
	go get ./...

vet: deps
	go vet ./...
