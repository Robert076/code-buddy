package lint_go

import (
	"fmt"

	"github.com/Robert076/devops-buddy/internal/command"
)

type LintGoCommand struct {
}

func (l *LintGoCommand) Name() string {
	return "go"
}

func (l *LintGoCommand) Description() string {
	return "Returns a Github Action with a go linter"
}

func (l *LintGoCommand) Run(args []string) error {
	// should be replaced with an api call in the future
	fmt.Println("name: linter\non:\n  push:\n    branches: [\"main\"]\n    paths: [\"**.go\"]\n  pull_request:\n    branches: [\"main\"]\n    paths: [\"**.go\"]\njobs:\n  lint:\n    name: Lint\n    runs-on: ubuntu-latest\n    timeout-minutes: 3\n    steps:\n      - uses: actions/checkout@v4\n        with:\n          fetch-depth: 1\n      - name: Setup Go\n        uses: actions/setup-go@v5\n        with:\n          go-version: \"1.23\"\n      - name: Install golangci-lint\n        run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest\n      - name: Run golangci-lint\n        run: golangci-lint run ./cmd/devops-buddy")
	return nil
}

func (l *LintGoCommand) Subcommands() []command.Command {
	return nil // leaf command
}
