package lint_command

import "github.com/Robert076/devops-buddy/internal/command"

type LintCommand struct{}

func (l *LintCommand) Name() string {
	return "lint"
}

func (l *LintCommand) Description() string {
	return "Returns linters for supported languages"
}

func (l *LintCommand) Run(args []string) error {
	return nil
}

func (l *LintCommand) Subcommands() []command.Command {
	return []command.Command{}
}
