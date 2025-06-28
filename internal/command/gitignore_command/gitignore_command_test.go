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

func TestGitignoreSubcommand_Name(t *testing.T) {
	cmd := &GitignoreSubcommand{"Go"}
	want := "Go"

	if got := cmd.Name(); got != want {
		t.Fatalf("GitignoreSubcommand_Name() got %s, want %s", got, want)
	}
}

func TestGitignoreSubcommand_Description(t *testing.T) {
	cmd := &GitignoreSubcommand{"Go"}
	want := "Returns a gitignore template for Go projects"

	if got := cmd.Description(); got != want {
		t.Fatalf("GitignoreSubcommand_Description() got %s, want %s", got, want)
	}
}

func TestGitignoreSubcommand_Subcommands(t *testing.T) {
	cmd := &GitignoreSubcommand{"Go"}
	subcommands := cmd.Subcommands()

	if len(subcommands) != 0 {
		t.Fatalf("Expected length of subcommands 0, got: %d", len(subcommands))
	}
}

func TestGitignoreSubcommand_Run(t *testing.T) {
	tests := []struct {
		name        string
		cmd         *GitignoreSubcommand
		expectErr   bool
		errMsg      string
		description string
	}{
		{
			"Valid - integration test, tests external endpoint",
			&GitignoreSubcommand{"this-gitignore-template-doesnt-exist"},
			true,
			"bad return code fetching github repo for this-gitignore-template-doesnt-exist gitignore: 404",
			"User requests a gitignore for an unknown template by running 'gitignore this-gitignore-template-doesnt-exist', this will cause a 404 return code (not found)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cmd.Run([]string{})

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
