package cmd

import "github.com/urfave/cli/v2"

//CmdPulls is test command.
var CmdPulls = cli.Command{
	Name:        "pulls",
	Aliases:     []string{"pull", "pr"},
	Usage:       "Manage and checkout pull requests",
	Description: `Manage and checkout pull requests`,
	ArgsUsage:   "[<pull index>]",
	Action:      runEmpty,
}

func runEmpty(ctx *cli.Context) error {
	return nil
}
