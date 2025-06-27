package main

import (
	"fmt"

	"github.com/Robert076/code-buddy/internal/registry"
)

func main() {
	var reg = registry.NewCommandRegistry()

	reg.InitRegistry()

	var args []string
	args = append(args, "go")
	err := reg.RunCommand("lint", args)
	if err != nil {
		fmt.Print("Hello World!")
	}
}
