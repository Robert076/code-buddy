package gitignore_command

import "testing"

func TestGitignoreCommand_Name(t *testing.T) {
	cmd := &GitignoreCommand{}
	want := "gitignore"
	if got := cmd.Name(); got != want {
		t.Fatalf("GitignoreCommand.Name() = %v, want %v", got, want)
	}
}

func TestGitignoreCommand_Description(t *testing.T) {
	cmd := &GitignoreCommand{}
	want := "Returns gitignore templates from the official Github documentation"
	if got := cmd.Description(); got != want {
		t.Fatalf("GitignoreCommand.Description() = %v, want %v", got, want)
	}
}

func TestGitignoreCommand_Run(t *testing.T) {
	tests := []struct {
		name string

		args    []string
		wantErr bool
	}{}
}
func TestGitignoreCommand_Subcommands(t *testing.T) {
	cmd := &GitignoreCommand{}
	subcommands := cmd.Subcommands()

	if len(subcommands) != 1 {
		t.Errorf("Expected %d subcommands, got %d", 1, len(subcommands))
	}
}
