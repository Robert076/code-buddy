package capture_output

import (
	"bytes"
	"os"
)

func CaptureOutput(f func()) string {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = origStdout
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
