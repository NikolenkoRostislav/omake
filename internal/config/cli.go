package config

import (
	"fmt"
	"os"

	"github.com/rostislav/omake/internal/state"
)

func ShowConfigPath() {
	path, err := getConfigPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Configuration file path:", path)
}

func SetupConfig() {
	var path string
	var err error

	if len(os.Args) < 4 {
		fmt.Println("Custom config path not provided, using default config path")
		path = expandHome("~\\omake")
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating default config directory:", err)
			return
		}
	} else {
		path = os.Args[3]
		path, err = FindDir(path)
		if err != nil {
			fmt.Println("Error finding directory:", err)
			return
		}
	}

	if err := createConfigFiles(path); err != nil {
		fmt.Println("Error creating config files:", err)
		return
	}

	state, err := state.Load()
	if err != nil {
		fmt.Println("Error loading state:", err)
		return
	}

	state.ConfigPath = path
	if err := state.Save(); err != nil {
		fmt.Println("Error saving state:", err)
		return
	}

	fmt.Println("Configuration file path:", path)
}
