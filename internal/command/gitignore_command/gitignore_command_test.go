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
	integrationTests := []struct {
		name        string
		args        []string
		expectErr   bool
		errMsg      string
		description string
	}{
		{
			"Valid - integration test, tests external endpoint",
			[]string{"go"},
			false,
			"",
			"User requests a gitignore for go by running 'gitignore go', this is an integration test it is not mocked",
		},
		{
			"Invalid - wrong argument given",
			[]string{"this-gitignore-template-doesnt-exist"},
			true,
			"unknown gitignore template: this-gitignore-template-doesnt-exist",
			"User requests a gitignore template that doesn't exist",
		},
	}

	for _, tt := range integrationTests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &GitignoreCommand{}
			err := cmd.Run(tt.args)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
				if err.Error() != tt.errMsg {
					t.Fatalf("Unexpected error: got %s, wanted %s", err, tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got '%v'", err)
				}
			}
		})
	}
}
func TestGitignoreCommand_Subcommands(t *testing.T) {
	cmd := &GitignoreCommand{}
	subcommands := cmd.Subcommands()

	if len(subcommands) != 155 {
		t.Errorf("Expected %d subcommands, got %d", 155, len(subcommands))
	}
}
