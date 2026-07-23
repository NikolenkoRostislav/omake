package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rostislav/omake/internal/state"
)

func GetConfigPath() {
	path, err := getConfigPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Configuration file path:", path)
}

func GetMakefilePath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	makefilePath := filepath.Join(configPath, "Makefile")
	if _, err := os.Stat(makefilePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Makefile does not exist in the configuration directory")
		return "", errors.New("Makefile does not exist in the configuration directory")
	}

	return makefilePath, nil
}

func GetYamlConfigPath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	makefilePath := filepath.Join(configPath, "config.yaml")
	if _, err := os.Stat(makefilePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("config.yaml does not exist in the configuration directory")
		return "", errors.New("config.yaml does not exist in the configuration directory")
	}

	return makefilePath, nil
}

func SetupConfig() {
	var path string
	if len(os.Args) < 4 {
		fmt.Println("Custom config path not provided, using default config path")
		path = expandHome("~\\omake")
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating default config directory:", err)
			return
		}
	} else {
		path = os.Args[3]
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

func getConfigPath() (string, error) {
	state, err := state.Load()
	if err != nil {
		return "", errors.New("Error loading state")
	}

	path := state.ConfigPath
	if path == "" {
		return "", errors.New("Configuration file path not set")
	}

	return path, nil
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
