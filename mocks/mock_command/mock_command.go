package mock_command

import "github.com/Robert076/devops-buddy/internal/command"

type MockCommand struct {
	name        string
	description string
	expectedErr error
	called      bool
	args        []string
	subcommands []command.Command
}

func (m *MockCommand) Name() string        { return m.name }
func (m *MockCommand) Description() string { return m.description }
func (m *MockCommand) Run(args []string) error {
	m.called = true
	m.args = args
	return m.expectedErr
}
func (m *MockCommand) Subcommands() []command.Command { return m.subcommands }
