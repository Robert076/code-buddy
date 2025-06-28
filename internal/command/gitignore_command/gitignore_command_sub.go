package gitignore_command

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Robert076/code-buddy/internal/command"
	"github.com/Robert076/code-buddy/internal/constants"
	"github.com/Robert076/code-buddy/internal/utils/format_json"
)

type GitignoreSubcommand struct {
	Language string
}

func (g *GitignoreSubcommand) Name() string {
	return g.Language
}

func (g *GitignoreSubcommand) Description() string {
	return "Returns a gitignore template for " + g.Language + " projects"
}

func (g *GitignoreSubcommand) Run(args []string) error {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(constants.GithubGitignoreLink + g.Language)
	if err != nil {
		return fmt.Errorf("error fetching github repo for %s gitignore: %v", g.Language, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad return code fetching github repo for %s gitignore: %d", g.Language, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body for %s gitignore: %v", g.Language, err)
	}

	formatted := format_json.FormatNiceStringFromUglyJSON(body)
	fmt.Println(formatted)

	return nil
}

func (g *GitignoreSubcommand) Subcommands() []command.Command {
	return nil
}

func NewGitignoreSubcommand(language string) command.Command {
	return &GitignoreSubcommand{Language: language}
}
