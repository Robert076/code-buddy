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
