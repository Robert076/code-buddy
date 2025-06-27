package main

import (
	"fmt"
	"os"

	"github.com/Robert076/code-buddy/internal/registry"
)

func main() {
	var reg = registry.NewCommandRegistry()

	reg.InitRegistry()

	err := reg.RunCommand(os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Printf("Error running command: %v", err)
	}
}
