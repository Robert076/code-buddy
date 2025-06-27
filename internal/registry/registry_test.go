package registry

import (
	"testing"
)

func TestCommandRegistry_New(t *testing.T) {
	cmd := NewCommandRegistry()
	if cmd == nil {
		t.Fatal("Expected non-nil CommandRegistry")
	}
}

func TestCommandRegistry_Init(t *testing.T) {
	cmd := NewCommandRegistry()

	cmd.InitRegistry()

	if cmd.commands == nil {
		t.Fatal("Expected non-nil list of commands in registry")
	}

	if len(cmd.commands) != 1 {
		t.Fatalf("Expected length of commands %d, got %d", 1, len(cmd.commands))
	}
}

func TestCommandRegistry_Run(t *testing.T) {
	tests := []struct {
		name        string
		command     string
		args        []string
		expectErr   bool
		errMsg      string
		description string
	}{
		{
			"Valid - run 'lint go' command",
			"lint",
			[]string{"go"},
			false,
			"",
			"User runs the 'lint go' command",
		},
		{
			"Invalid - unknown command (command not recognized)",
			"lintt",
			[]string{"go"},
			true,
			"command 'lintt' is not recognised",
			"User runs 'lint' command for a linter that doesnt exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommandRegistry()

			cmd.InitRegistry()

			err := cmd.RunCommand(tt.command, tt.args)

			if tt.expectErr {
				if err == nil {
					t.Fatal("Expected error, got none")
				} else if err.Error() != tt.errMsg {
					t.Fatalf("Expected error '%v', got '%v'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Didn't expect error, got '%v'", err.Error())
				}
			}
		})
	}
}
