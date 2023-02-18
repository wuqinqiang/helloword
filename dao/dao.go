package dao

type Dao struct {
	Word       Word
	Phrase     Phrase
	WordPhrase WordPhrase
}

func Get() Dao {
	return Dao{
		Word:       NewWord(),
		Phrase:     NewPhrase(),
		WordPhrase: NewWordPhrase(),
	}
}
