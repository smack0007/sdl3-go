package main

import (
	"fmt"

	SDL "github.com/smack0007/sdl-go/sdl"
)

func main() {
	fmt.Printf("=== Compiled ===\n")
	fmt.Printf("Version: %d.%d.%d\n", SDL.VERSIONNUM_MAJOR(SDL.VERSION), SDL.VERSIONNUM_MINOR(SDL.VERSION), SDL.VERSIONNUM_MICRO(SDL.VERSION))
	fmt.Printf("Revision: %s\n", SDL.REVISION)

	fmt.Printf("\n")

	fmt.Printf("=== Linked ===\n")
	version := SDL.GetVersion()
	fmt.Printf("Version: %d.%d.%d\n", SDL.VERSIONNUM_MAJOR(version), SDL.VERSIONNUM_MINOR(version), SDL.VERSIONNUM_MICRO(version))
	fmt.Printf("Revision: %s\n", SDL.GetRevision())
}
