// Copyright 2020 The Christmastree Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Christmastree is command line tool for christmastree
package main // import "code.gitea.io/christmastree"

import (
	"fmt"
	"os"
	"strings"

	"github.com/khmarbaise/christmastree/cmd"

	"github.com/urfave/cli/v2"
)

// Version holds the current tea version
var Version = "development"

// Tags holds the build tags used
var Tags = ""

func main() {
	app := cli.NewApp()
	app.Name = "Christmastree"
	app.Usage = "Command line tool to interact with Christmastree"
	app.Version = Version + formatBuiltWith(Tags)
	app.Commands = []*cli.Command{
		&cmd.CmdTree,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}
}

func formatBuiltWith(Tags string) string {
	if len(Tags) == 0 {
		return ""
	}

	return " built with: " + strings.Replace(Tags, " ", ", ", -1)
}
