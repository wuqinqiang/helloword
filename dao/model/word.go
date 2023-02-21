package model

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Words []*Word

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

func (words Words) ListByLetter() map[string]Words {
	m := make(map[string]Words)
	for _, item := range words {
		if len(item.Word) == 0 {
			continue
		}
		startLetter := item.Word[0:1]
		m[startLetter] = append(m[startLetter], item)
	}
	return m
}

func (words Words) RandomPick() *Word {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
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
