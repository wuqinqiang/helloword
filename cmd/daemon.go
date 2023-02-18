package cmd

import (
	"os"

	"github.com/wuqinqiang/helloword/collector/bbdc"

	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/helloword/collector"
	"github.com/wuqinqiang/helloword/conf"
	"github.com/wuqinqiang/helloword/core"
	"github.com/wuqinqiang/helloword/generator/gpt3"
	"github.com/wuqinqiang/helloword/notify"
	"github.com/wuqinqiang/helloword/selector"
)

var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "daemon",
	Action: func(context *cli.Context) error {
		settings, err := conf.GetConf()
		if err != nil {
			return err
		}
		generator := gpt3.NewClient(settings.GptToken)

		var collectors []collector.Collector
		bbcdCookie := os.Getenv("BBDC_COOKIE")
		if bbcdCookie != "" {
			collectors = append(collectors, bbdc.New(bbcdCookie))
		}
		importer := collector.NewImporter(collectors...)

		s := selector.New(selector.Random)
		n := notify.New(settings.Senders())
		core := core.New(generator, importer, s, n)
		return core.Run()
	},
}
