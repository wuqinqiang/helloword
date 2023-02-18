package selector

import (
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
	NextWords() (words model.Words, err error)
	SetStrategyType(strategy StrategyType)
}
type Strategy interface {
	Select(words model.Words) model.Words
}

type Srv struct {
	wordDao   dao.Word
	maxNumber int
	strategy  Strategy
}

func New(strategyType StrategyType, wordDao dao.Word) Selector {
	srv := &Srv{
		wordDao:   wordDao,
		maxNumber: 6,
	}
	var strategy Strategy

	switch strategyType {
	case Default:
	case LeastRecentlyUsed:
	default:
		strategy = s.NewRandom(srv.maxNumber)
	}

	srv.strategy = strategy

	return srv
}

func (s Srv) NextWords() (words model.Words, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Srv) SetStrategyType(strategy StrategyType) {
	//TODO implement me
	panic("implement me")
}
