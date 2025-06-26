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

type MockCommandBuilder struct {
	mock *MockCommand
}

func NewMockCommandBuilder() *MockCommandBuilder {
	return &MockCommandBuilder{
		mock: &MockCommand{},
	}
}

func (b *MockCommandBuilder) WithName(name string) *MockCommandBuilder {
	b.mock.name = name
	return b
}

func (b *MockCommandBuilder) WithDescription(description string) *MockCommandBuilder {
	b.mock.description = description
	return b
}

func (b *MockCommandBuilder) WithSubcommands(subcommands []command.Command) *MockCommandBuilder {
	b.mock.subcommands = subcommands
	return b
}

func (b *MockCommandBuilder) Build() *MockCommand {
	return b.mock
}
