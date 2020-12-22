package cmd

import (
	"fmt"
	"github.com/khmarbaise/christmastree/modules"
	"github.com/urfave/cli/v2"
)

//CmdTree is test command.
var CmdTree = cli.Command{
	Name:        "tree",
	Aliases:     []string{"tr"},
	Usage:       "create a tree",
	Description: `Create a tree`,
	Action:      runEmpty,
}

func runEmpty(ctx *cli.Context) error {
	fmt.Printf("%s\n", modules.TreeOutput(5))
	return nil
}
