build:
	go build -o ./bin/sdl-go ./sdl

clean:
	go clean
	rm ./bin/sdl-go

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