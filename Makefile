REPO_PATH := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

CP    := cp -u
MKDIR := mkdir -p

build:
	go build -o ./bin/sdl-go/sdl ./sdl
	go build -o ./bin/sdl-go/img ./img

build-examples:
	go build -o ./bin/examples/template ./examples/template.go
	@for dir in ./examples/*/*; do echo "Building $$dir..."; go build -o ./bin/$$dir $$dir; done

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

run-example-img-01:
	$(MKDIR) ./examples/img/01-show-image/assets
	$(CP) ./assets/sample.png ./examples/img/01-show-image/assets/sample.png
	go run ./examples/img/01-show-image

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