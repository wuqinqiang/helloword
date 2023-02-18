package dao

import (
	"context"

	"github.com/wuqinqiang/helloword/dao/model"
)

type Phrase interface {
	Create(ctx context.Context, phrase *model.Phrase) error
}

type PhraseImpl struct {
}

func NewPhrase() Phrase {
	return &PhraseImpl{}
}

func (impl PhraseImpl) Create(ctx context.Context, phrase *model.Phrase) error {
	w := use(ctx).Phrase
	return w.WithContext(ctx).Create(phrase)
}
