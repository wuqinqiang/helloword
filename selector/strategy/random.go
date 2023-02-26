package strategy

import (
	"math/rand"
	"time"

	"github.com/wuqinqiang/helloword/dao/model"
)

type RandomStrategy struct{}

func NewRandom() *RandomStrategy {
	return &RandomStrategy{}
}

func (s *RandomStrategy) Select(words model.Words) model.Words {
	shuffled := make(model.Words, len(words))
	copy(shuffled, words)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
