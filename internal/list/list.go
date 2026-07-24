package list

import (
	"fmt"

	"github.com/rostislav/omake/internal/config"
)

func ShowList() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	for target := range config.Targets {
		fmt.Println(target)
	}
}
