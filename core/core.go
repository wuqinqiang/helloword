package core

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/wuqinqiang/helloword/notify/base"

	"github.com/wuqinqiang/helloword/logging"

	"github.com/wuqinqiang/helloword/selector"

	"github.com/wuqinqiang/helloword/tools"

	"github.com/wuqinqiang/helloword/collector"

	"github.com/robfig/cron/v3"
	"github.com/wuqinqiang/helloword/generator"
	"github.com/wuqinqiang/helloword/notify"
)

type Core struct {
	generator generator.Generator
	notify    notify.Notify
	cron      *cron.Cron
	options   *Options
	importer  *collector.Importer
	selector  selector.Selector
	ch        chan struct{}

	locker sync.Mutex
}

func New(generator generator.Generator, importer *collector.Importer,
	selector selector.Selector, notify notify.Notify, opts ...Option) *Core {
	core := &Core{
		generator: generator,
		notify:    notify,
		options:   Default,
		cron:      cron.New(),
		importer:  importer,
		selector:  selector,
		ch:        make(chan struct{}),
	}

	for _, opt := range opts {
		opt(core.options)
	}

	return core
}

func (core *Core) Run() error {
	defer core.cron.Stop()

	tools.GoSafe(func() {
		core.runImport()
	})

	// generatePhrase
	if _, err := core.cron.AddFunc(core.options.spec, core.generatePhrase); err != nil {
		return err
	}
	core.cron.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
	return nil
}

func (core *Core) runImport() {
	_ = core.importer.Import(context.Background())
}

func (core *Core) generatePhrase() {
	defer core.locker.Unlock()
	core.locker.Lock()
	words, err := core.selector.NextWords()
	if err != nil {
		logging.Errorf("[NextWords] err:%v", err)
		return
	}
	if len(words) == 0 {
		logging.Warnf("no words available")
		return
	}
	phrase, err := core.generator.Generate(context.Background(), words.List())
	if err != nil {
		return
	}
	// todo records for phrase
	core.notify.Notify(base.New("", phrase))
}
