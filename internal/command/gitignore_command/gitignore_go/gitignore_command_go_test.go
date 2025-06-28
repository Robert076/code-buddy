package gitignore_go

import (
	"testing"

	"github.com/Robert076/code-buddy/internal/utils/capture_output"
	"github.com/stretchr/testify/assert"
)

func TestGitignoreCommandGo_Name(t *testing.T) {
	cmd := &GitignoreCommandGo{}
	want := "go"

	if got := cmd.Name(); got != want {
		t.Errorf("GitignoreCommandGo_Name() got %v, want %v", got, want)
	}
}

func TestGitignoreCommandGo_Description(t *testing.T) {
	cmd := &GitignoreCommandGo{}
	want := "Returns a gitignore template for go projects"

	if got := cmd.Description(); got != want {
		t.Errorf("GitignoreCommandGo_Description() got %v, want %v", got, want)
	}
}

func TestGitignoreCommandGo_Subcommands(t *testing.T) {
	cmd := &GitignoreCommandGo{}
	noSubcommands := len(cmd.Subcommands())

	if noSubcommands != 0 {
		t.Errorf("Expected subcommands length 0, since it's a leaf command, got %v", noSubcommands)
	}
}

func TestGitignoreCommandGo_Run(t *testing.T) {
	cmd := &GitignoreCommandGo{}

	output := capture_output.CaptureOutput(func() {
		cmd.Run([]string{})
	})
	assert.Contains(t, output, ".env")
	assert.Contains(t, output, "# env file")
	assert.Contains(t, output, ".exe")
}
