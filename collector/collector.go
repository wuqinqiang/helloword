package collector

import (
	"context"
	"sync"

	"github.com/wuqinqiang/helloword/dao"

	"github.com/wuqinqiang/helloword/logging"

	"github.com/wuqinqiang/helloword/dao/model"
)

type Collector interface {
	Name() string
	Collect(ctx context.Context) (model.Words, error)
}

type Importer struct {
	wordDao    dao.Word
	collectors []Collector
}

func NewImporter(wordDao dao.Word, collectors ...Collector) *Importer {
	return &Importer{
		wordDao:    wordDao,
		collectors: collectors,
	}
}

func (importer Importer) Import(ctx context.Context) error {
	var wg sync.WaitGroup

	for i := range importer.collectors {
		wg.Add(1)

		go func(collector Collector) {
			defer wg.Done()
			words, err := collector.Collect(ctx)
			if err != nil {
				logging.Errorf("[Import] collect:%s err:%v", collector.Name(), err)
				return
			}

			if err = importer.wordDao.BatchInsert(ctx, words); err != nil {
				logging.Errorf("[Import] BatchInsert err:%v", err)
			}

		}(importer.collectors[i])
	}

	wg.Wait()
	return nil
}
