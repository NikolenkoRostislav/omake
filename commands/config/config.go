package config

import (
	"fmt"
	"os"
)

func ShowConfig() {
	if len(os.Args) < 3 {
		fmt.Println("omake: missing configuration file path")
		return
	}
	fmt.Println("Configuration file path:", os.Args[2])
}
