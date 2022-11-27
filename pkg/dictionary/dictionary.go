package dictionary

// Word lists are read assuming that all words are the same WordType.
// A word is registered in the Dictionary with the highest WordType
// value, i.e. if a word exists only in a Bonus list it is Bonus
// type, if it exists in Bonus and Normal it is Normal type, and
// if it exists at all in a Swear list it is always a Swear word.
type WordType int

const (
	NoWord WordType = iota
	BonusWord
	NormalWord
	SwearWord
)

type HasWordsWithPrefixResult struct {
	HasThisWord        WordType
	HasWordsWithPrefix bool
}

type WordReader interface {
	ReadWord() string
}

type PrefixDictionary interface {
	HasWord(string) *HasWordsWithPrefixResult
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

func (dictionary *prefixDictionary) Read(reader WordReader, wordTypes WordType, upgradeOnly bool) {
	for {
		word := reader.ReadWord()
		if word == "" {
			break
		}
		for i := 1; i < len(word); i++ {
			prefix := word[0:i]
			if v, ok := dictionary.Prefixes[prefix]; ok {
				v.HasWordsWithPrefix = true
			} else if !upgradeOnly {
				dictionary.Prefixes[prefix] = &HasWordsWithPrefixResult{
					HasThisWord:        NoWord,
					HasWordsWithPrefix: true,
				}
			}
		}
		if v, ok := dictionary.Prefixes[word]; ok {
			if wordTypes > v.HasThisWord {
				v.HasThisWord = wordTypes
			}
		} else if !upgradeOnly {
			dictionary.Prefixes[word] = &HasWordsWithPrefixResult{
				HasThisWord:        wordTypes,
				HasWordsWithPrefix: false,
			}
		}
	}
}

func (p *prefixDictionary) HasWord(prefix string) *HasWordsWithPrefixResult {
	if res, ok := p.Prefixes[prefix]; ok {
		return res
	} else {
		return &HasWordsWithPrefixResult{}
	}
}
