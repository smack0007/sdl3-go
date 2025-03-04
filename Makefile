build:
	go build -o ./bin/sdl-go ./sdl && echo "Build succeeded."

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