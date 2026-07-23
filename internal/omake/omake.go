package omake

import (
	"fmt"
	"os"

	"github.com/rostislav/omake/internal/config"
)

func getMakefileAndConfig() (string, string, error) {
	configFile, err := config.GetYamlConfigPath()
	if err != nil {
		return "", "", err
	}

	makefileFile, err := config.GetMakefilePath()
	if err != nil {
		return "", "", err
	}

	configData, err := os.ReadFile(configFile)
	if err != nil {
		return "", "", err
	}

	makefileData, err := os.ReadFile(makefileFile)
	if err != nil {
		return "", "", err
	}

	return string(configData), string(makefileData), nil
}

func Make(command string) {
	config, makefile, err := getMakefileAndConfig()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(config, makefile)
	// TODO: Implement
}
