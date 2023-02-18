package model

import (
	"time"

	"github.com/google/uuid"
)

func NewPhrase(phrase string) *Phrase {
	now := time.Now().Unix()
	return &Phrase{
		PhraseID:   "p" + uuid.NewString(),
		Phrase:     phrase,
		CreateTime: now,
		UpdateTime: now,
	}
}
