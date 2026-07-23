package help

func ShowHelp() {
	println("You can use the following commands:")
	println("  help - Show this help message")
	println("  version - Show the version information")
	println("  config path - Show configuration file path")
	println("  config setup - Initialize configuration")
	println("  config setup <path> - Initialize configuration with a custom path")
	println("  <target> - Run the specified target from the Makefile")
}
