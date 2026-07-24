package omake

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rostislav/omake/internal/config"
)

func Make(command string) {
	makefilePath, err := config.GetMakefilePath()
	if err != nil {
		fmt.Println("Error getting Makefile path:", err)
		return
	}

	config, err := config.GetConfig()
	if err != nil {
		fmt.Println("Error getting configuration:", err)
		return
	}

	executionDir, err := chooseExecutionDirectory(config)
	if err != nil {
		fmt.Println("Error choosing execution directory:", err)
		return
	}

	execMakeCommand(command, executionDir, makefilePath)
}

func chooseExecutionDirectory(config *config.Config) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return cwd, nil
}

func execMakeCommand(command string, executionDir string, makefilePath string) {
	cmd := exec.Command("make", "-f", makefilePath, "-C", executionDir, command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
