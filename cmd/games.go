package cmd

import (
	"github.com/urfave/cli/v2"
)

var GamesCmd = &cli.Command{
	Name:  "games",
	Usage: "import your own words",
	Subcommands: []*cli.Command{
		wordChainCmd,
	},
}
