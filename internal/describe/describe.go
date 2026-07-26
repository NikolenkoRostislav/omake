package describe

import (
	"errors"
	"fmt"

	"github.com/rostislav/omake/internal/config"
)

func ShowDescription(target string) {
	targetConfig, err := config.GetConfigForTarget(target)
	if err != nil {
		if errors.Is(err, config.ErrTargetNotFound) {
			fmt.Println("Target not found.")
			return
		}
		panic(err)
	}

	if targetConfig.Description == "" {
		fmt.Println("No description available for this target.")
	} else {
		fmt.Println(targetConfig.Description)
	}

	if len(targetConfig.Variables) > 0 {
		fmt.Println("\nVariables:")
		for _, variable := range targetConfig.Variables {
			varDescription := variable.Description
			fmt.Printf("  %s: %s\n", variable.Name, varDescription)

			varEnv := variable.EnvVar
			if varEnv != "" {
				fmt.Printf("    Environment variable: %s\n", varEnv)
			} else {
				fmt.Println("    Warning: No environment variable specified for this variable.")
			}

			varDefault := variable.Default
			if varDefault != "" {
				fmt.Printf("    Default: %s\n", varDefault)
			}
		}
	}
}
