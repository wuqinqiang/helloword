package cmd

import (
	"os"

	"github.com/wuqinqiang/helloword/collector/file"

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
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "files",
			Usage: "传入要导入的Library目录下的单词文件。例如你可以导入一个文件 damon --files=CET4.txt" +
				"或者多个文件用逗号隔开，damon --files=CET4.txt,CET6.txt",
		},
		&cli.StringFlag{
			Name: "spec",
			Usage: "推送时间频率控制,默认1小时推送一次短语。自定义比如每5分钟推送一次: @every 5m。" +
				"其他规则参考库github.com/robfig/cron",
		},
	},

	Action: func(cctx *cli.Context) error {
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

		files := "CET4.txt"
		if cctx.String("files") != "" {
			files = cctx.String("files")
		}
		collectors = append(collectors, file.New(files))

		importer := collector.NewImporter(collectors...)

		s := selector.New(selector.Random)
		n := notify.New(settings.Senders())
		core := core.New(generator, importer, s, n, core.WithSpec(cctx.String("spec")))

		return core.Run()
	},
}
