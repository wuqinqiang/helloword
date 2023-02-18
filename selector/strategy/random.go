package strategy

import (
	"math/rand"
	"time"

	"github.com/wuqinqiang/helloword/dao/model"
)

type RandomStrategy struct {
	maxWords int
}

func NewRandom(maxWords int) *RandomStrategy {
	return &RandomStrategy{
		maxWords: maxWords,
	}
}

func (s *RandomStrategy) Select(words model.Words) model.Words {
	if len(words) <= s.maxWords {
		return words
	}
	shuffled := make(model.Words, len(words))
	copy(shuffled, words)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled[:s.maxWords]
}
