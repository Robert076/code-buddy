package lint_command

import (
	"testing"
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
	if len(subs) != 1 {
		t.Errorf("Expected one subcommand in list, got %d", len(subs))
	}
}

func TestLintCommand_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		expectErr  bool
		errMsg     string
		descripton string
	}{
		{
			"Valid - no arguments provided, should print instructions",
			[]string{},
			false,
			"",
			"User runs the Lint command without providing arguments. It should print instructions.",
		},
		{
			"Invalid - unknown linter",
			[]string{"this-linter-doesnt-exist"},
			true,
			"unknown linter: this-linter-doesnt-exist",
			"User runs the Lint command and specifies a linter that is NOT supported. It should throw an error.",
		},
		{
			"Valid - known linter",
			[]string{"go"},
			false,
			"",
			"User runs the Lint command and specifies a linter that is supported. It should work fine.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &LintCommand{}
			err := cmd.Run(tt.args)

			if tt.expectErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				if err.Error() != tt.errMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got '%v'", err)
				}
			}
		})
	}
}
