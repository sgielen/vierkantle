package dictionary

type HasWordsWithPrefixResult struct {
	HasThisWord        bool
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

func NewPrefixDictionaryFromFile(file string, minPrefixLength, maxPrefixLength int) (PrefixDictionary, error) {
	reader, err := NewFileReader(file)
	if err != nil {
		return nil, err
	}
	return NewPrefixDictionary(reader, minPrefixLength, maxPrefixLength), nil
}

func NewPrefixDictionary(reader WordReader, minPrefixLength, maxPrefixLength int) PrefixDictionary {
	dictionary := &prefixDictionary{
		Prefixes: map[string]*HasWordsWithPrefixResult{},
	}

	// TODO: do something with minPrefixLength and maxPrefixLength
	for {
		word := reader.ReadWord()
		if word == "" {
			break
		}
		for i := 1; i < len(word); i++ {
			prefix := word[0:i]
			if v, ok := dictionary.Prefixes[prefix]; ok {
				v.HasWordsWithPrefix = true
			} else {
				dictionary.Prefixes[prefix] = &HasWordsWithPrefixResult{
					HasThisWord:        false,
					HasWordsWithPrefix: true,
				}
			}
		}
		if v, ok := dictionary.Prefixes[word]; ok {
			v.HasThisWord = true
		} else {
			dictionary.Prefixes[word] = &HasWordsWithPrefixResult{
				HasThisWord:        true,
				HasWordsWithPrefix: false,
			}
		}
	}

	return dictionary
}

func (p *prefixDictionary) HasWord(prefix string) *HasWordsWithPrefixResult {
	if res, ok := p.Prefixes[prefix]; ok {
		return res
	} else {
		return &HasWordsWithPrefixResult{}
	}
}
