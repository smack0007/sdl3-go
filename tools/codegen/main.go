package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to get cwd", err)
		os.Exit(1)
	}

	sdlJsonBytes, err := os.ReadFile(path.Join(cwd, "sdl.json"))

	if err != nil {
		log.Fatal("Failed to load sdl.json", err)
		os.Exit(1)
	}

	var sdlJson []FFIJsonNode
	json.Unmarshal(sdlJsonBytes, &sdlJson)

	for _, node := range sdlJson {
		if !strings.HasPrefix(node.Location, "/usr/local/include/SDL3") {
			continue
		}

		fmt.Printf("%#v\n", node)
	}
}
