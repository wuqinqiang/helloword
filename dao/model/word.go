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

func (word *Word) SetDefinition(definition string) {
	word.Definition = definition
}

func (word *Word) SetPhonetic(phonetic string) {
	word.Phonetic = phonetic
}

func (words Words) List() (items []string) {
	for i := range words {
		items = append(items, words[i].Word)
	}
	return
}
func (words Words) WordIds() (items []string) {
	for i := range words {
		items = append(items, words[i].WordID)
	}
	return
}

func (words Words) GenerateWordPhrase(phraseId string) (list []*WordPhrase) {
	for _, word := range words {
		list = append(list, NewWordPhrase(word.WordID, phraseId))
	}
	return list
}
