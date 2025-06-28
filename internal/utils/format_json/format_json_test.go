package format_json

import (
	"testing"
)

func TestFormatJSON(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		want        string
		description string
	}{
		{
			name:        "Valid JSON - should extract source field",
			body:        "{\"name\":\"Go\",\"source\":\"line1\\nline2\\nline3\"}",
			want:        "line1\nline2\nline3",
			description: "Expected to extract the 'source' field from valid JSON",
		},
		{
			name:        "Invalid JSON - should return input as-is",
			body:        "name\":\"Go\",\"source\":\"broken json",
			want:        "name\":\"Go\",\"source\":\"broken json",
			description: "Expected to return original input if JSON is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatNiceStringFromUglyJSON([]byte(tt.body))
			if got != tt.want {
				t.Errorf("FAILED: %s\nGot:\n%q\nWant:\n%q", tt.description, got, tt.want)
			}
		})
	}
}
