package main

import (
	"fmt"

	SDL "github.com/smack0007/sdl_go/sdl"
)

func main() {

	var version SDL.Version
	SDL.GetVersion(&version)
	fmt.Printf("SDL Version: %d.%d.%d\n", version.Major, version.Minor, version.Patch)

	fmt.Printf("SDL Revision: %s\n", SDL.GetRevision())
}
