package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rostislav/omake/internal/state"
)

func GetMakefilePath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	makefilePath := filepath.Join(configPath, "Makefile")
	if _, err := os.Stat(makefilePath); errors.Is(err, os.ErrNotExist) {
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
		return "", errors.New("config.yaml does not exist in the configuration directory")
	}

	return makefilePath, nil
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

func findDir(path string) (string, error) {
	path = filepath.Clean(path)
	path = expandHome(path)

	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}
	if !info.IsDir() {
		return "", errors.New("path is a file, not a directory")
	}

	return path, nil
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
