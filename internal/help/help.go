package help

import "fmt"

func ShowHelp() {
	fmt.Println("You can use the following commands:")
	fmt.Println("  help - Show this help message")
	fmt.Println("  version - Show the version information")
	fmt.Println("  config path - Show configuration file path")
	fmt.Println("  config setup - Initialize configuration")
	fmt.Println("  config setup <path> - Initialize configuration with a custom path")
	fmt.Println("  describe <target> - Show description and variables for the specified target")
	fmt.Println("  list - List all targets present in config.yaml")
	fmt.Println("  <target> - Run the specified target from the Makefile")
}
