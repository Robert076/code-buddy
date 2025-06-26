package lint_go

import (
	"testing"

	capture_output "github.com/Robert076/devops-buddy/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestLintCommandGo_Name(t *testing.T) {
	cmd := &LintGoCommand{}
	if got := cmd.Name(); got != "go" {
		t.Errorf("LintCommandGo.Name() = %v, want %v", got, "go")
	}
}

func TestLintCommandGo_Description(t *testing.T) {
	cmd := &LintGoCommand{}
	if got := cmd.Description(); got != "Returns a Github Action with a go linter" {
		t.Errorf("LintCommandGo.Description() = %v, want %v", got, "Returns a Github Action with a go linter")
	}
}

func TestLintCommandGo_Run(t *testing.T) {
	cmd := &LintGoCommand{}
	output := capture_output.CaptureOutput(func() {
		cmd.Run([]string{})
	})
	assert.Contains(t, output, "name: linter")
	assert.Contains(t, output, "push:")
	assert.Contains(t, output, "golangci-lint run")
}

func TestLintCommandGo_Subcommands(t *testing.T) {
	cmd := &LintGoCommand{}
	subs := cmd.Subcommands()
	if len(subs) != 0 { // this is a leaf command
		t.Errorf("LintCommandGo.Subcommands() returned non-empty subcommands list, want 0, got %v", len(subs))
	}
}
