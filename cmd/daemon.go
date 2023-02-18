package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/helloword/collector"
	"github.com/wuqinqiang/helloword/conf"
	"github.com/wuqinqiang/helloword/core"
	"github.com/wuqinqiang/helloword/dao"
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

		dao := dao.WordImpl{}
		importer := collector.NewImporter(dao)
		s := selector.New(selector.Random, dao)

		n := notify.New(settings.Senders())
		core := core.New(generator, importer, s, n)
		return core.Run()
	},
}
