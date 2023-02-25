package cmd

import (
	"errors"
	"strings"

	"github.com/wuqinqiang/helloword/logging"

	"github.com/wuqinqiang/helloword/notify/base"

	"github.com/wuqinqiang/helloword/notify"

	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/helloword/conf"
	"github.com/wuqinqiang/helloword/generator/gpt3"
)

var PhraseCmd = &cli.Command{
	Name:  "phrase",
	Usage: "Generate phrases directly",
	Action: func(cctx *cli.Context) error {
		req := cctx.Args().Get(0)
		if req == "" {
			return errors.New("please input your words")
		}
		conf, err := conf.GetConf()
		if err != nil {
			return err
		}
		client := gpt3.NewClient(conf.GptToken)
		phrase, err := client.Generate(cctx.Context, strings.Split(req, ","))
		if err != nil {
			return err
		}
		n := notify.New(conf.Senders())
		n.Notify(base.New("", phrase))

		n.Wait()
		logging.Infof("Successfully generated a short phrase")
		return nil
	},
}
