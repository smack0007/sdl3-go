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
	go run ./examples/hello-world/main.go

run-scraper:
	go run ./tools/scraper/main.go $(shell pkg-config sdl3 --cflags)

run-sdlinfo:
	go run ./tools/sdlinfo/main.go