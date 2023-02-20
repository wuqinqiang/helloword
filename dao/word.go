package dao

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"github.com/wuqinqiang/helloword/dao/model"
)

type Word interface {
	BatchInsert(ctx context.Context, words model.Words) error
	GetList(ctx context.Context) (words model.Words, err error)
	IncrNumRepetitions(ctx context.Context, wordIds []string) error
}

type WordImpl struct {
}

func NewWord() Word {
	return &WordImpl{}
}

func (impl WordImpl) BatchInsert(ctx context.Context, words model.Words) error {
	w := use(ctx).Word
	return w.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "word"}},
		DoUpdates: clause.AssignmentColumns([]string{"phonetic", "definition", "difficulty"}),
	}).CreateInBatches(words, 50)
}

func (impl WordImpl) IncrNumRepetitions(ctx context.Context, wordIds []string) error {
	w := use(ctx).Word
	now := time.Now().Unix()
	_, err := w.WithContext(ctx).Where(w.WordID.In(wordIds...)).
		UpdateSimple(w.NumRepetitions.Add(1), w.LastUsed.Value(now))
	return err
}

func (impl WordImpl) GetList(ctx context.Context) (words model.Words, err error) {
	w := use(ctx).Word
	words, err = w.WithContext(ctx).Limit(limit).Find()
	return
}
