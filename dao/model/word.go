package model

import (
	"time"

	"github.com/google/uuid"
)

type Words []*Word

func NewWords(items []string) (words Words) {
	for _, item := range items {
		words = append(words, NewWord(item))
	}
	return
}

func NewWord(word string) *Word {
	now := time.Now().Unix()
	return &Word{
		WordID:     uuid.NewString(),
		Word:       word,
		CreateTime: now,
		UpdateTime: now,
	}
}
