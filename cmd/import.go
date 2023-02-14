package cmd

import "github.com/urfave/cli/v2"

var ImportCmd = &cli.Command{
	Name:  "import",
	Usage: "Batch import of English words",
	Action: func(context *cli.Context) error {
		return nil
	},
}
