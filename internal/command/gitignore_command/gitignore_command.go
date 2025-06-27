package gitignore_command

import (
	"fmt"
	"strings"

	"github.com/Robert076/code-buddy/internal/command"
	"github.com/Robert076/code-buddy/internal/command/gitignore_command/gitignore_go"
)

type GitignoreCommand struct {
}

func (g *GitignoreCommand) Name() string {
	return "gitignore"
}

func (g *GitignoreCommand) Description() string {
	return "Returns gitignore templates from the official Github documentation"
}

func (g *GitignoreCommand) Run(args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: do gitignore <language>")
		fmt.Println("Supported gitignore templates:")
		for _, sub := range g.Subcommands() {
			fmt.Printf("  %s - %s\n", sub.Name(), sub.Description())
		}
		return nil
	}

	for _, sub := range g.Subcommands() {
		if strings.EqualFold(sub.Name(), args[0]) {
			return sub.Run(args[1:])
		}
	}

	return fmt.Errorf("unknown gitignore template: %s", args[0])
}

func (g *GitignoreCommand) Subcommands() []command.Command {
	return []command.Command{
		&gitignore_go.GitignoreCommandGo{},
	}
}
