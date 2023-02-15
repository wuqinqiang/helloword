package cmd

import "github.com/urfave/cli/v2"

var GameCmd = &cli.Command{
	Name:        "game",
	Usage:       "daemon",
	Subcommands: []*cli.Command{},
}
