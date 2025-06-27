package registry

import (
	"fmt"

	"github.com/Robert076/code-buddy/internal/command"
	"github.com/Robert076/code-buddy/internal/command/gitignore_command"
	"github.com/Robert076/code-buddy/internal/command/lint_command"
)

type CommandRegistry struct {
	commands map[string]command.Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]command.Command),
	}
}

func (c *CommandRegistry) RegisterCommand(cmd command.Command) {
	c.commands[cmd.Name()] = cmd
}

func (c *CommandRegistry) InitRegistry() {
	c.RegisterCommand(&lint_command.LintCommand{})
	c.RegisterCommand(&gitignore_command.GitignoreCommand{})
}

func (c *CommandRegistry) RunCommand(name string, args []string) error {
	cmd, ok := c.commands[name]
	if ok {
		err := cmd.Run(args)
		return err
	} else {
		return fmt.Errorf("command '%s' is not recognised", name)
	}
}
