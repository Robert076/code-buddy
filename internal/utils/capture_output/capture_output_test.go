package capture_output

import (
	"testing"

	"github.com/Robert076/code-buddy/internal/command/gitignore_command/gitignore_go"
	"github.com/stretchr/testify/assert"
)

func TestCaptureOutput(t *testing.T) {
	cmd := &gitignore_go.GitignoreCommandGo{}
	output := CaptureOutput(func() {
		cmd.Run([]string{})
	})
	assert.Contains(t, output, ".env")
	assert.Contains(t, output, "# env file")
	assert.Contains(t, output, ".exe")
}
