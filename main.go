package main

import (
	"fmt"
	"os"

	"github.com/rostislav/omake/internal/config"
	"github.com/rostislav/omake/internal/help"
	"github.com/rostislav/omake/internal/omake"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing command, use 'omake help' for usage information")
		return
	}

	switch os.Args[1] {
	case "help":
		help.ShowHelp()
	case "version":
		fmt.Println("omake v0.1.0")
	case "config":
		if len(os.Args) < 3 {
			fmt.Println("config command requires an additional argument, use 'omake help' for usage information")
			return
		}

		switch os.Args[2] {
		case "path":
			config.ShowConfigPath()
		case "setup":
			config.SetupConfig()

		default:
			fmt.Println("Unknown config command, use 'omake help' for usage information")
		}

	default:
		omake.Make(os.Args[1])
	}
}
