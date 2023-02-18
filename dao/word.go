package dao

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/wuqinqiang/helloword/dao/model"
)

type Word interface {
	BatchInsert(ctx context.Context, words model.Words) error
}

type WordImpl struct {
}

func (impl WordImpl) BatchInsert(ctx context.Context, words model.Words) error {
	w := use(ctx).Word
	return w.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "word"}},
		DoUpdates: clause.AssignmentColumns([]string{"phonetic", "definition", "difficulty", "update_time"}),
	}).CreateInBatches(words, 50)
}

func (impl WordImpl) GetList(ctx context.Context) (words model.Words, err error) {
	w := use(ctx).Word
	words, err = w.WithContext(ctx).Limit(limit).Find()
	return
}
