package model

import (
	"time"

	"github.com/google/uuid"
)

func NewWordPhrase(wordId, phraseId string) *WordPhrase {
	now := time.Now().Unix()
	return &WordPhrase{
		WordPhraseID: "wp" + uuid.NewString(),
		WordID:       wordId,
		PhraseID:     phraseId,
		CreateTime:   now,
		UpdateTime:   now,
	}
}
