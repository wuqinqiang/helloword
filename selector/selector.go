package selector

import (
	"context"

	"github.com/wuqinqiang/helloword/dao"
	"github.com/wuqinqiang/helloword/dao/model"
	s "github.com/wuqinqiang/helloword/selector/strategy"
)

type StrategyType string

const (
	Default           StrategyType = "default"
	Random            StrategyType = "random"
	LeastRecentlyUsed StrategyType = "leastRecentlyUsed"
)

type Selector interface {
	NextWords(ctx context.Context) (words model.Words, err error)
	SetStrategyType(strategy StrategyType)
}
type Strategy interface {
	Select(words model.Words) model.Words
}

type Srv struct {
	dao.Dao
	maxNumber int
	strategy  Strategy
}

func New(strategyType StrategyType) Selector {
	srv := &Srv{
		Dao:       dao.Get(),
		maxNumber: 6,
	}
	var strategy Strategy

	switch strategyType {
	case LeastRecentlyUsed:
	case Default:
	default:
		strategy = s.NewRandom(srv.maxNumber)
	}

	srv.strategy = strategy

	return srv
}

func (s *Srv) NextWords(ctx context.Context) (words model.Words, err error) {
	words, err = s.Word.GetList(ctx)
	if err != nil {
		return
	}
	words = s.strategy.Select(words)
	return
}

func (s *Srv) SetStrategyType(strategy StrategyType) {
	//todo
}
