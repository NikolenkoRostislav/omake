package main

import (
	"fmt"
	"os"

	"github.com/rostislav/omake/commands/config"
	"github.com/rostislav/omake/commands/help"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("omake: missing command, use -h or --help for usage information")
		return
	}

	switch os.Args[1] {
	case "-h", "--help":
		help.ShowHelp()
	case "-v", "--version":
		fmt.Println("omake v0.1.0")
	case "-cfg", "--config":
		config.ShowConfig()
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
