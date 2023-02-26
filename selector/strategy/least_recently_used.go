package strategy

import (
	"sort"

	"github.com/wuqinqiang/helloword/dao/model"
)

type LeastRecentlyUsed struct{}

func NewLeastRecentlyUsed() *LeastRecentlyUsed {
	return &LeastRecentlyUsed{}
}

func (l LeastRecentlyUsed) Select(words model.Words) model.Words {
	// sort the words by NumRepetitions field in ascending order
	sort.Slice(words, func(i, j int) bool {
		return words[i].NumRepetitions < words[j].NumRepetitions
	})

	// sort the words by last_used field in ascending order
	sort.Slice(words, func(i, j int) bool {
		return words[i].LastUsed < words[j].LastUsed
	})
	return words
}
