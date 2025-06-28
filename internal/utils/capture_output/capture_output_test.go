package capture_output

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureOutput(t *testing.T) {
	output := CaptureOutput(func() {
		fmt.Print("Hello world")
	})
	assert.Contains(t, output, "Hello")
	assert.Contains(t, output, "world")
}
