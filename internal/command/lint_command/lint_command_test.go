package lint_command

import (
	"testing"

	"github.com/Robert076/devops-buddy/mocks/mock_command"
)

func TestLintCommand_Name(t *testing.T) {
	cmd := &LintCommand{}
	if got := cmd.Name(); got != "lint" {
		t.Errorf("LintCommand.Name() = %v, want %v", got, "lint")
	}
}

func TestLintCommand_Description(t *testing.T) {
	cmd := &LintCommand{}
	if got := cmd.Description(); got != "Returns linters for supported languages" {
		t.Errorf("LintCommand.Description() = %v, want %v", got, "Returns linters for supported languages")
	}
}

func TestLintCommand_Subcommands(t *testing.T) {
	cmd := &LintCommand{}
	subs := cmd.Subcommands()
	if len(subs) != 0 {
		t.Errorf("Expected empty subcommands list, got %d", len(subs))
	}
}

func TestLintCommand_Run(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		subcommands []*mock_command.MockCommand
		wantErr     bool
		errMsg      string
		description string
	}{
		{
			"Invalid - no arguments provided",
			[]string{},
			[]*mock_command.MockCommand{},
			false,
			"",
			"When no args provided, should print usage and return nil",
		},
		{
			"Invalid - unsupported linter",
			[]string{"go"},
			[]*mock_command.MockCommand{
				mock_command.NewMockCommandBuilder().WithName("Python").WithDescription("Python linter").Build(),
			},
			true,
			"unknown linter: go",
			"When a linter that is not supported is provided, we should get an error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &LintCommand{}
			err := cmd.Run(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				} else if err.Error() != tt.errMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}
