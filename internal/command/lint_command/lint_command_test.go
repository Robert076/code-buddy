package lint_command

import "testing"

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
