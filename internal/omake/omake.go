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

	execMakeCommand(target, executionDir, makefilePath)
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

func execMakeCommand(target string, executionDir string, makefilePath string) {
	cmd := exec.Command("make", "-f", makefilePath, "-C", executionDir, target)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
