package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Targets map[string]Target `yaml:"targets"`
}

type Target struct {
	ExecutionDir string              `yaml:"execution_dir"`
	Description  string              `yaml:"description"`
	Variables    map[string]Variable `yaml:"variables"`
}

type Variable struct {
	Description string `yaml:"description"`
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
