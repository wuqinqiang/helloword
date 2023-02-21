package games

import (
	"testing"

	"github.com/wuqinqiang/helloword/dao/model"
)

func TestWordChain(t *testing.T) {
	words := model.Words{
		model.NewWord("apple"),
		model.NewWord("youngster"),
		model.NewWord("ear"),
		model.NewWord("art"),
		model.NewWord("eraser"),
		model.NewWord("rival"),
		model.NewWord("luxury"),
	}

	startWord := model.NewWord("apple")
	chain := NewWordChain(words, startWord)

	// Test if startWord is set correctly
	if chain.StartWord().Word != startWord.Word {
		t.Errorf("Start word is not set correctly")
	}

	// Test if the first pick is correct
	nextWord, ok := chain.Pick()
	if !ok {
		t.Errorf("Failed to pick next word")
	} else if !(nextWord.Word == "ear" || nextWord.Word == "eraser") {
		t.Errorf("Unexpected word picked")
	}

	// Test if the second pick is correct
	nextWord, ok = chain.Pick()
	if !ok {
		t.Errorf("Failed to pick next word")
	} else if nextWord.Word != "rival" {
		t.Errorf("Unexpected word picked")
	}

	// Test if the third pick is correct
	nextWord, ok = chain.Pick()
	if !ok {
		t.Errorf("Failed to pick next word")
	} else if nextWord.Word != "luxury" {
		t.Errorf("Unexpected word picked")
	}

	// Test if the fourth pick is correct
	nextWord, ok = chain.Pick()
	if !ok {
		t.Errorf("Failed to pick next word")
	} else if nextWord.Word != "youngster" {
		t.Errorf("Unexpected word picked")
	}

	// Test if the last pick returns false as expected
	// rival was used
	_, ok = chain.Pick()
	if ok {
		t.Errorf("Expected last pick to return false")
	}
}
