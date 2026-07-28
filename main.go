package main

import (
	"fmt"
	"os"

	"github.com/rostislav/omake/internal/config"
	"github.com/rostislav/omake/internal/describe"
	"github.com/rostislav/omake/internal/help"
	"github.com/rostislav/omake/internal/list"
	"github.com/rostislav/omake/internal/omake"
)

func main() {
	if len(os.Args) < 2 {
		help.ShowHelp()
		return
	}

	switch os.Args[1] {
	case "help":
		help.ShowHelp()
	case "version":
		fmt.Println("omake v0.1.0")
	case "cfg-path":
		config.ShowConfigPath()
	case "init":
		if len(os.Args) < 3 {
			config.SetupConfig("")
		} else {
			config.SetupConfig(os.Args[2])
		}
	case "desc":
		if len(os.Args) < 3 {
			fmt.Println("desc command requires an additional argument, use 'omake help' for usage information")
			return
		}
		describe.ShowDescription(os.Args[2])
	case "list":
		list.ShowList()

	default:
		omake.Make(os.Args[1])
	}
}
