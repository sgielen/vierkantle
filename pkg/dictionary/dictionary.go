package dictionary

import "strings"

// Word lists are read assuming that all words are the same WordType.
// A word is registered in the Dictionary with the highest WordType
// value, i.e. if a word exists only in a Bonus list it is Bonus
// type, if it exists in Bonus and Normal it is Normal type, and
// if it exists at all in a Sensitive list it is always a Sensitive word.
type WordType int

const (
	NoWord WordType = iota
	BonusWord
	NormalWord
	SensitiveWord
)

type HasWordsWithPrefixResult struct {
	HasThisWord        WordType
	Frequency          float64
	HasWordsWithPrefix bool
}

type WordReadResult struct {
	Word      string
	Frequency float64
}

type WordReader interface {
	ReadWord() WordReadResult
}

type PrefixDictionary interface {
	HasWord(string) *HasWordsWithPrefixResult
}

type RWPrefixDictionary interface {
	HasWord(string) *HasWordsWithPrefixResult
	AddWord(WordReadResult, WordType, bool)
}

type prefixDictionary struct {
	Prefixes map[string]*HasWordsWithPrefixResult
}

func NewPrefixDictionary() *prefixDictionary {
	return &prefixDictionary{
		Prefixes: map[string]*HasWordsWithPrefixResult{},
	}
}

func (dictionary *prefixDictionary) ReadFromFile(file string, wordTypes WordType, upgradeOnly bool) error {
	reader, err := NewFileReader(file)
	if err != nil {
		return err
	}
	dictionary.Read(reader, wordTypes, upgradeOnly)
	return nil
}

func (dictionary *prefixDictionary) AddWord(word WordReadResult, wordType WordType, upgradeOnly bool) {
	word.Word = NormalizeWord(word.Word)
	for i := 1; i < len(word.Word); i++ {
		prefix := word.Word[0:i]
		if v, ok := dictionary.Prefixes[prefix]; ok {
			v.HasWordsWithPrefix = true
		} else if !upgradeOnly {
			dictionary.Prefixes[prefix] = &HasWordsWithPrefixResult{
				HasThisWord:        NoWord,
				Frequency:          0,
				HasWordsWithPrefix: true,
			}
		}
	}
	if v, ok := dictionary.Prefixes[word.Word]; ok {
		if wordType > v.HasThisWord {
			v.HasThisWord = wordType
		}
		if word.Frequency > v.Frequency {
			v.Frequency = word.Frequency
		}
	} else if !upgradeOnly {
		dictionary.Prefixes[word.Word] = &HasWordsWithPrefixResult{
			HasThisWord:        wordType,
			Frequency:          word.Frequency,
			HasWordsWithPrefix: false,
		}
	}
}

func (dictionary *prefixDictionary) Read(reader WordReader, wordTypes WordType, upgradeOnly bool) {
	for {
		word := reader.ReadWord()
		if word.Word == "" {
			break
		}
		dictionary.AddWord(word, wordTypes, upgradeOnly)
	}
}

func (p *prefixDictionary) HasWord(prefix string) *HasWordsWithPrefixResult {
	if res, ok := p.Prefixes[prefix]; ok {
		return res
	} else {
		return &HasWordsWithPrefixResult{}
	}
}

func NormalizeWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	return word
}
