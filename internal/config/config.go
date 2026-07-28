package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Targets map[string]Target `yaml:"targets"`
}

type Target struct {
	ExecutionDir string     `yaml:"exec_dir"`
	Description  string     `yaml:"desc"`
	Variables    []Variable `yaml:"vars"`
}

type Variable struct {
	Name        string `yaml:"name"`
	Description string `yaml:"desc"`
	EnvVar      string `yaml:"env_var"`
	Default     string `yaml:"default"`
}

func GetConfig() (*Config, error) {
	path, err := GetYamlConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

var ErrTargetNotFound = errors.New("target not found")

func GetConfigForTarget(targetName string) (*Target, error) {
	cfg, err := GetConfig()
	if err != nil {
		return nil, err
	}

	target, ok := cfg.Targets[targetName]
	if !ok {
		return nil, ErrTargetNotFound
	}

	return &target, nil
}
