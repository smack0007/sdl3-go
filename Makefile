REPO_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	go build -o ./bin/sdl-go ./sdl

build-examples:
	@find ./examples/*/* -type d -exec sh -c 'echo "Building {}..." && go build -o ./bin/{} {}' ';'

clean:
	go clean
	rm -rf ./bin

lint:
	golangci-lint run ./...

vet:
	go vet ./...

run-hello-world:
	go run ./examples/hello-world

run-scraper:
	go run ./tools/scraper $(shell pkg-config sdl3 --cflags)

run-sdlinfo:
	go run ./tools/sdlinfo

run-example-renderer-clear:
	go run ./examples/renderer/01-clear

run-example-renderer-primitives:
	go run ./examples/renderer/02-primitives

run-example-renderer-lines:
	go run ./examples/renderer/03-lines

run-example-renderer-points:
	go run ./examples/renderer/04-points

run-example-renderer-rectangles:
	go run ./examples/renderer/05-rectangles

run-example-renderer-textures:
	go run ./examples/renderer/06-textures