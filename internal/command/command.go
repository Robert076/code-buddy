package command

type Command interface {
	Name() string
	Description() string
	Run(args []string) error
	Subcommands() []Command
}
