package state

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type State struct {
	ConfigPath string `json:"configPath"`
}

func (state State) Save() error {
	path, err := getStatePath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Load() (State, error) {
	var state State

	path, err := getStatePath()
	if err != nil {
		return state, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return state, state.Save()
		}

		return state, err
	}

	err = json.Unmarshal(data, &state)
	if err != nil {
		return state, err
	}

	return state, nil
}

func getStatePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	omakeDir := filepath.Join(configDir, "omake")

	return filepath.Join(omakeDir, "state.json"), nil
}
