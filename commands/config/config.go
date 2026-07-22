package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rostislav/omake/state"
)

func GetConfigPath() {
	state, err := state.Load()
	if err != nil {
		fmt.Println("Error loading state:", err)
		return
	}

	path := state.ConfigPath
	if path == "" {
		fmt.Println("Configuration file path not set")
		return
	}

	fmt.Println("Configuration file path:", path)
}

func SetupConfig() {
	var path string
	if len(os.Args) < 3 {
		fmt.Println("Custom config path not provided, using default config path")
		path = expandHome("~\\omake")
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating default config directory:", err)
			return
		}
	} else {
		path = os.Args[2]
		path = findDir(path)
		if path == "" {
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

func findDir(path string) string {
	path = filepath.Clean(path)
	path = expandHome(path)

	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error: cannot get absolute path:")
		return ""
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("Error: path does not exist")
		return ""
	}
	if !info.IsDir() {
		fmt.Println("Error: path is a file, not a directory")
		return ""
	}

	return path
}

func expandHome(path string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	if path == "~" {
		return home
	}

	if rest, ok := strings.CutPrefix(path, `~\`); ok {
		return filepath.Join(home, rest)
	}

	return path
}

func createConfigFiles(path string) error {
	makefilePath := filepath.Join(path, "Makefile")
	configPath := filepath.Join(path, "config.yaml")

	if err := createFileIfNotExists(makefilePath, ""); err != nil {
		return err
	}

	if err := createFileIfNotExists(configPath, ""); err != nil {
		return err
	}

	return nil
}

func createFileIfNotExists(path, content string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return os.WriteFile(path, []byte(content), 0644)
}
