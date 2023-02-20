package cmd

import "github.com/urfave/cli/v2"

// ImportCmd todo import your own words
var ImportCmd = &cli.Command{
	Name:  "import",
	Usage: "import your own words",
	Action: func(context *cli.Context) error {
		return nil
	},
}
