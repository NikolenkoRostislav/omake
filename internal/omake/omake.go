package omake

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/rostislav/omake/internal/config"
)

func Make(target string) {
	makefilePath, err := config.GetMakefilePath()
	if err != nil {
		fmt.Println("Error getting Makefile path:", err)
		return
	}

	cfg, err := config.GetConfigForTarget(target)
	if err != nil && !errors.Is(err, config.ErrTargetNotFound) {
		fmt.Println("Error getting configuration:", err)
		return
	}

	executionDir, err := chooseExecutionDirectory(cfg)
	if err != nil {
		fmt.Println("Error choosing execution directory:", err)
		return
	}

	variables, err := getVariables(cfg)
	if err != nil {
		fmt.Println("Error getting variables:", err)
		return
	}

	execMakeCommand(target, executionDir, makefilePath, variables)
}

func chooseExecutionDirectory(targetConfig *config.Target) (string, error) {
	if targetConfig == nil || targetConfig.ExecutionDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return cwd, nil
	}

	dir, err := config.FindDir(targetConfig.ExecutionDir)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func getVariables(targetConfig *config.Target) (map[string]string, error) {
	variables := make(map[string]string)

	if targetConfig == nil || len(targetConfig.Variables) == 0 {
		return variables, nil
	}

	for name, variable := range targetConfig.Variables {
		if variable.EnvVar == "" {
			fmt.Printf("Warning: No environment variable specified for variable '%s'. The value will be ignored.\n", name)
			continue
		}

		value, found := findArgumentValue(name)
		if !found {
			if variable.Default == "" {
				return nil, fmt.Errorf("required variable '%s' is not set", name)
			}

			value = variable.Default
		}

		variables[variable.EnvVar] = value
	}

	return variables, nil
}

func findArgumentValue(name string) (string, bool) {
	for i, arg := range os.Args {
		if arg == name {
			if i+1 >= len(os.Args) {
				fmt.Printf("Error: variable '%s' was provided without a value\n", name)
				return "", false
			}

			return os.Args[i+1], true
		}
	}

	return "", false
}

func execMakeCommand(target string, executionDir string, makefilePath string, variables map[string]string) {
	cmd := exec.Command("make", "-f", makefilePath, "-C", executionDir, target)

	cmd.Env = os.Environ()

	for key, value := range variables {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
