package registry

import (
	"github.com/Robert076/devops-buddy/internal/command"
	"github.com/Robert076/devops-buddy/internal/command/lint_command"
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
}

func (c *CommandRegistry) RunCommand(name string, args []string) error {
	c.commands[name].Run(args)
	return nil
}
