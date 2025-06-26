package lint_command

import (
	"fmt"

	"github.com/Robert076/devops-buddy/internal/command"
)

type LintCommand struct{}

func (l *LintCommand) Name() string {
	return "lint"
}

func (l *LintCommand) Description() string {
	return "Returns linters for supported languages"
}

func (l *LintCommand) Run(args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: do lint <language>")
		fmt.Println("Supported linters:")
		for _, sub := range l.Subcommands() {
			fmt.Printf("  %s - %s\n", sub.Name(), sub.Description())
		}
		return nil
	}

	for _, sub := range l.Subcommands() {
		if sub.Name() == args[0] {
			return sub.Run(args[1:])
		}
	}

	return fmt.Errorf("unknown linter: %s", args[0])
}

func (l *LintCommand) Subcommands() []command.Command {
	return []command.Command{}
}

type LintCommandBuilder struct {
	lint *LintCommand
}
