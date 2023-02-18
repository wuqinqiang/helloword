package core

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/wuqinqiang/helloword/dao"
	"github.com/wuqinqiang/helloword/dao/model"

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
	dao dao.Dao

	generator generator.Generator
	notify    notify.Notify
	cron      *cron.Cron
	importer  *collector.Importer
	selector  selector.Selector
	ch        chan struct{}
	locker    sync.Mutex

	options *Options
}

func New(generator generator.Generator, importer *collector.Importer,
	selector selector.Selector, notify notify.Notify, opts ...Option) *Core {
	core := &Core{
		dao:       dao.Get(),
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

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent,
		time.Duration(core.options.selectTimeout)*time.Second)
	defer cancel()

	words, err := core.selector.NextWords(ctx)
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
	tools.GoSafe(func() {
		core.notify.Notify(base.New("本次短语", phrase))
	})

	core.afterGenerate(phrase, words)
}
func (core *Core) afterGenerate(phrase string, words model.Words) {
	ctx := context.Background()
	phraseRecord := model.NewPhrase(phrase)
	if err := core.dao.Phrase.Create(ctx, phraseRecord); err != nil {
		logging.Errorf("Create Phrase err:%v", err)
		return
	}
	if err := core.dao.WordPhrase.BatchInsert(ctx,
		words.GenerateWordPhrase(phraseRecord.PhraseID)); err != nil {
		logging.Errorf("WordPhrase BatchInsert err:%v", err)
		return
	}
	if err := core.dao.Word.IncrNumRepetitions(ctx, words.WordIds()); err != nil {
		logging.Errorf("IncrNumRepetitions err:%v", err)
		return
	}
}
