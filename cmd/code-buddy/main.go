package main

import (
	"fmt"
	"os"

	"github.com/Robert076/code-buddy/internal/registry"
)

func main() {
	var reg = registry.NewCommandRegistry()

	reg.InitRegistry()

	if len(os.Args) == 1 {
		fmt.Printf(`=====================================================
Welcome to code-buddy. Your CLI friend.
Example command: 'code-buddy lint go'
=====================================================`)
	} else {
		err := reg.RunCommand(os.Args[1], os.Args[2:])
		if err != nil {
			fmt.Printf("Error running command: %v", err)
		}
	}
}
