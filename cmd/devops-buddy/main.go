package main

import (
	"github.com/Robert076/devops-buddy/internal/registry"
)

func main() {
	var reg = registry.NewCommandRegistry()

	reg.InitRegistry()

	var args []string
	args = append(args, "go")
	reg.RunCommand("lint", args)
}
