package dao

import (
	"context"

	"github.com/wuqinqiang/helloword/dao/model"
)

type WordPhrase interface {
	BatchInsert(ctx context.Context, list []*model.WordPhrase) error
}

type WordPhraseImpl struct {
}

func NewWordPhrase() WordPhrase {
	return &WordPhraseImpl{}
}

func (impl WordPhraseImpl) BatchInsert(ctx context.Context, list []*model.WordPhrase) error {
	w := use(ctx).WordPhrase
	return w.WithContext(ctx).CreateInBatches(list, 50)
}
