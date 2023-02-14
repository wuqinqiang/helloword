package main

import (
	"github.com/urfave/cli/v2"
	"helloword/cmd"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "hello word",
		Usage: "happy study english word",
		Before: func(context *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			cmd.DaemonCmd,
			cmd.ImportCmd,
		},
	}
	app.Setup()
	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString("Error:" + err.Error() + "\n")
	}
}
