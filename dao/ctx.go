package dao

import (
	"context"

	"github.com/wuqinqiang/helloword/dao/query"
	"github.com/wuqinqiang/helloword/db"
)

const (
	limit = 10000
)

func use(ctx context.Context) *query.Query {
	db := db.Get().WithContext(ctx)
	return query.Use(db)
}
