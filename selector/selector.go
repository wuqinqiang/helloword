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

type Option func(srv *Srv)

func WithWordNumber(wordNumer int) Option {
	return func(srv *Srv) {
		srv.wordNumber = wordNumer
	}
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
		srv.wordNumber = 5
	}
	if srv.wordNumber > 10 {
		srv.wordNumber = 10
	}

	var strategy Strategy

	switch strategyType {
	// todo
	case LeastRecentlyUsed:
	case Default:
	default:
		strategy = s.NewRandom(srv.wordNumber)
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
