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

run-sdlinfo:
	go run ./tools/sdlinfo/main.go