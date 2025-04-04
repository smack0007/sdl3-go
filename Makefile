REPO_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	go build -o ./bin/sdl-go ./sdl

build-examples:
	go build -o ./bin/examples/template ./examples/template.go
	@find ./examples/*/* -type d -exec sh -c 'echo "Building {}..." && go build -o ./bin/{} {}' ';'

clean:
	go clean
	rm -rf ./bin

lint:
	golangci-lint run ./...

vet:
	go vet ./...

run-scraper:
	go run ./tools/scraper $(shell pkg-config sdl3 --cflags)

run-sdlinfo:
	go run ./tools/sdlinfo

run-example-pollevent-01:
	go run ./examples/pollevent/01-basic

run-example-renderer-01:
	go run ./examples/renderer/01-clear

run-example-renderer-02:
	go run ./examples/renderer/02-primitives

run-example-renderer-03:
	go run ./examples/renderer/03-lines

run-example-renderer-04:
	go run ./examples/renderer/04-points

run-example-renderer-05:
	go run ./examples/renderer/05-rectangles

run-example-renderer-06:
	go run ./examples/renderer/06-textures

run-example-renderer-07:
	go run ./examples/renderer/07-streaming-textures

run-example-renderer-08:
	go run ./examples/renderer/08-rotating-textures