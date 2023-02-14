package cmd

import "github.com/urfave/cli/v2"

var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "daemon",
	Action: func(context *cli.Context) error {
		return nil
	},
}
