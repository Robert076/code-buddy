package gitignore_go

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Robert076/code-buddy/internal/command"
	"github.com/Robert076/code-buddy/internal/constants"
	"github.com/Robert076/code-buddy/internal/utils/format_json"
)

type GitignoreCommandGo struct {
}

func (g *GitignoreCommandGo) Name() string {
	return "go"
}

func (g *GitignoreCommandGo) Description() string {
	return "Returns a gitignore template for go projects"
}

func (g *GitignoreCommandGo) Run(args []string) error {
	resp, err := http.Get(constants.GithubGitignoreLink)

	if err != nil {
		return fmt.Errorf("error fetching github repo for go gitignore: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad return code fetching github repo for go gitignore: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("error reading response body for go gitignore: %v", err)
	}

	formatted := format_json.FormatNiceStringFromUglyJSON(body)
	fmt.Println(formatted)

	return nil
}

func (g *GitignoreCommandGo) Subcommands() []command.Command {
	return nil // leaf command
}
