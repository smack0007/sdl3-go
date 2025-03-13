package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get cwd", err)
		os.Exit(1)
	}

	sdlJsonBytes, err := os.ReadFile(filepath.Join(cwd, "sdl.json"))
	if err != nil {
		log.Fatal("Failed to load sdl.json", err)
		os.Exit(1)
	}

	var sdlJson []FFIJsonNode
	json.Unmarshal(sdlJsonBytes, &sdlJson)

	sdlFileNamesToNodes := make(map[string][]FFIJsonNode)
	sdlImageNodes := make([]FFIJsonNode, 0)

	for _, node := range sdlJson {
		if strings.HasPrefix(node.Location, "/SDL3/") {
			locationParts := strings.Split(node.Location[len("/SDL3/"):], ":")
			fileName := locationParts[0]

			if sdlFileNamesToNodes[fileName] == nil {
				sdlFileNamesToNodes[fileName] = make([]FFIJsonNode, 0)
			}

			sdlFileNamesToNodes[fileName] = append(sdlFileNamesToNodes[fileName], node)
		} else if strings.HasPrefix(node.Location, "/SDL3_image/") {
			sdlImageNodes = append(sdlImageNodes, node)
		}
	}

	for fileName, nodes := range sdlFileNamesToNodes {
		outputFileName := fileName[len("SDL_"):len(fileName)-len(".h")] + ".go"
		outputFilePath := filepath.Join(cwd, "..", "..", "sdl", outputFileName)

		if outputFileName == "init.go" {
			generateFile(outputFilePath, nodes)
		}
	}
}

func generateFile(outputFilePath string, nodes []FFIJsonNode) {
	fmt.Println(outputFilePath)
	for _, node := range nodes {
		fmt.Println(node.Location, node.Tag, node.Name)
	}
}
