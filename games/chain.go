package games

import (
	"github.com/wuqinqiang/helloword/dao/model"
)

type WordChain struct {
	// dataSets [letter]->words  etc. [a]{apple,apply,application}
	dataSets map[string]model.Words
	// save words that playing games
	tmp map[string]struct{}
	//max tries for pick the word
	maxRetries int

	startWord *model.Word

	prevWord *model.Word
}

func NewWordChain(words model.Words, startWord *model.Word) *WordChain {
	wc := &WordChain{
		dataSets:   words.ListByLetter(),
		tmp:        make(map[string]struct{}),
		startWord:  startWord,
		prevWord:   startWord,
		maxRetries: 5,
	}
	wc.tmp[startWord.Word] = struct{}{}
	return wc
}

func (chain *WordChain) SetPrevWord(word string) bool {
	_, ok := chain.tmp[word]
	if ok {
		return false
	}
	chain.prevWord = model.NewWord(word)
	chain.tmp[word] = struct{}{}

	return true
}

func (chain *WordChain) StartWord() *model.Word {
	return chain.startWord
}

func (chain *WordChain) PrevWord() *model.Word {
	return chain.prevWord
}

func (chain *WordChain) Pick() (*model.Word, bool) {
	letterWords, ok := chain.dataSets[chain.prevWord.Word[len(chain.prevWord.Word)-1:]]
	if !ok {
		return nil, false
	}

	for i := 0; i < chain.maxRetries; i++ {
		w := letterWords.RandomPick()

		if chain.SetPrevWord(w.Word) {
			return w, true
		}
	}

	return nil, false
}
