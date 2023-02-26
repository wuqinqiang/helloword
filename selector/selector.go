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

	MinWordsNum = 5
	MaxWordsNum = 10
)

type Selector interface {
	NextWords(ctx context.Context) (words model.Words, err error)
	SetStrategyType(strategy Strategy)
}
type Strategy interface {
	Select(words model.Words) model.Words
}

type Srv struct {
	dao.Dao
	wordNumber int
	strategy   Strategy
}

func New(strategyType StrategyType, options ...Option) Selector {
	srv := &Srv{
		Dao: dao.Get(),
	}
	for _, option := range options {
		option(srv)
	}

	if srv.wordNumber <= 0 {
		srv.wordNumber = MinWordsNum
	}
	if srv.wordNumber > MaxWordsNum {
		srv.wordNumber = MaxWordsNum
	}

	var strategy Strategy
	switch strategyType {
	case LeastRecentlyUsed:
		strategy = s.NewLeastRecentlyUsed()
	case Default:
	default:
		strategy = s.NewRandom()
	}

	srv.strategy = strategy
	return srv
}

func (s *Srv) NextWords(ctx context.Context) (words model.Words, err error) {
	var list model.Words
	list, err = s.Word.GetList(ctx)
	if err != nil {
		return
	}

	words = s.strategy.Select(list)
	// not enough, then all out
	if len(words) < s.wordNumber {
		return
	}
	words = words[:s.wordNumber]

	return
}

func (s *Srv) SetStrategyType(strategy Strategy) {
	s.strategy = strategy
}
