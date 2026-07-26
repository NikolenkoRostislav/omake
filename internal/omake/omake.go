package omake

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

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
	if targetConfig == nil || len(targetConfig.Variables) == 0 {
		return nil, nil
	}

	if err := validateVariablesConfig(targetConfig.Variables); err != nil {
		return nil, err
	}

	args := os.Args[2:]
	providedVariables, err := getProvidedVariables(args)
	if err != nil {
		return nil, err
	}

	return getVariablesFromProvided(targetConfig.Variables, providedVariables)
}

func validateVariablesConfig(variables []config.Variable) error {
	seenWithDefault := false
	for _, variable := range variables {
		if variable.Name == "" {
			return errors.New("Variable name not specified")
		}
		if variable.EnvVar == "" {
			return fmt.Errorf("Environment variable not specified for variable %s", variable.Name)
		}
		if variable.Default == "" && seenWithDefault {
			return errors.New("Variables with default not allowed before variables without default")
		}
		if variable.Default != "" {
			seenWithDefault = true
		}
	}
	return nil
}

func getProvidedVariables(args []string) (map[string]string, error) {
	providedVariables := make(map[string]string)
	seenNamed := false
	for i, variable := range args {
		if strings.Contains(variable, "=") {
			seenNamed = true
			parts := strings.SplitN(variable, "=", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid variable format: %s", variable)
			}
			providedVariables[parts[0]] = parts[1]
			continue
		}

		if !seenNamed {
			providedVariables[strconv.Itoa(i)] = variable
		} else {
			return nil, errors.New("Positional variables can't be provided after keyword variables")
		}
	}

	return providedVariables, nil
}

func getVariablesFromProvided(neededVars []config.Variable, providedVars map[string]string) (map[string]string, error) {
	variables := make(map[string]string)

	for i, variable := range neededVars {
		if value, ok := providedVars[variable.Name]; ok {
			variables[variable.EnvVar] = value
		} else if value, ok := providedVars[strconv.Itoa(i)]; ok {
			variables[variable.EnvVar] = value
		} else if variable.Default != "" {
			variables[variable.EnvVar] = variable.Default
		} else {
			return nil, fmt.Errorf("Required variable '%s' is not set", variable.Name)
		}
	}

	return variables, nil
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
