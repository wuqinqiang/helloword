package main

import (
	"os"
	"strings"

	"github.com/wuqinqiang/helloword/db"

	"github.com/wuqinqiang/helloword/db/sqlite"

	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/helloword/cmd"
)

func main() {
	app := &cli.App{
		Name:  "hello word",
		Usage: "happy study english word",
		Before: func(context *cli.Context) error {
			dbPath := strings.TrimSpace(os.Getenv("HELLO_WORD_PATH"))
			if err := db.Init(sqlite.New(dbPath)); err != nil {
				return err
			}
			return nil
		},
		Commands: []*cli.Command{
			cmd.DaemonCmd,
			cmd.ImportCmd,
			cmd.GenCmd,
			cmd.PhraseCmd,
			cmd.GamesCmd,
		},
	}
	app.Setup()
	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString("Error:" + err.Error() + "\n")
	}
}
