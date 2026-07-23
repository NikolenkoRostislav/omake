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

	executionDir, err := chooseExecutionDirectory()
	if err != nil {
		fmt.Println("Error choosing execution directory:", err)
		return
	}

	cmd := exec.Command("make", "-f", makefilePath, "-C", executionDir, command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func chooseExecutionDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return cwd, nil
}
