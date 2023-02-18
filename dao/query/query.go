// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		Phrase:          newPhrase(db, opts...),
		Word:            newWord(db, opts...),
		WordPhrase:      newWordPhrase(db, opts...),
		WordPhraseUsage: newWordPhraseUsage(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Phrase          phrase
	Word            word
	WordPhrase      wordPhrase
	WordPhraseUsage wordPhraseUsage
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Phrase:          q.Phrase.clone(db),
		Word:            q.Word.clone(db),
		WordPhrase:      q.WordPhrase.clone(db),
		WordPhraseUsage: q.WordPhraseUsage.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Phrase:          q.Phrase.replaceDB(db),
		Word:            q.Word.replaceDB(db),
		WordPhrase:      q.WordPhrase.replaceDB(db),
		WordPhraseUsage: q.WordPhraseUsage.replaceDB(db),
	}
}

type queryCtx struct {
	Phrase          *phraseDo
	Word            *wordDo
	WordPhrase      *wordPhraseDo
	WordPhraseUsage *wordPhraseUsageDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Phrase:          q.Phrase.WithContext(ctx),
		Word:            q.Word.WithContext(ctx),
		WordPhrase:      q.WordPhrase.WithContext(ctx),
		WordPhraseUsage: q.WordPhraseUsage.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
