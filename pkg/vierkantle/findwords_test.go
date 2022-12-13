package vierkantle_test

import (
	"os"
	"sort"
	"testing"

	"github.com/go-test/deep"
	"github.com/sgielen/vierkantle/pkg/dictionary"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
)

func TestWordForPath(t *testing.T) {
	board := vierkantle.NewBoard(3, 3)
	board.Cells[0] = []rune{'h', 'e', 'x'}
	board.Cells[1] = []rune{'x', 'e', 'l'}
	board.Cells[2] = []rune{'x', 'o', 'l'}

	word := board.WordForPath(vierkantle.Path{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 1, Y: 2}})
	if word != "hello" {
		t.Errorf("Word should be 'hello', but was '%s'", word)
	}
}

type MockDictionary struct {
	ExistingPrefixes []string
	ExistingWords    []string
}

func (m *MockDictionary) HasWord(w string) *dictionary.HasWordsWithPrefixResult {
	res := dictionary.HasWordsWithPrefixResult{}
	for _, p := range m.ExistingWords {
		if p == w {
			res.HasThisWord = dictionary.NormalWord
		}
	}
	for _, p := range m.ExistingPrefixes {
		if p == w {
			res.HasWordsWithPrefix = true
		}
	}
	return &res
}

func TestWordsInBoard(t *testing.T) {
	board := vierkantle.NewBoard(3, 3)
	board.Cells[0] = []rune{'h', 'e', 'x'}
	board.Cells[1] = []rune{'x', 'i', 'l'}
	board.Cells[2] = []rune{'x', 'o', 'l'}

	mockDictionary := &MockDictionary{}
	mockDictionary.ExistingWords = []string{"hell", "hello", "hellish"}
	mockDictionary.ExistingPrefixes = []string{"h", "he", "hel", "hell", "hello", "helli", "hellis", "hellish"}

	wordsInBoard := board.WordsInBoard(mockDictionary, 4)
	var words []string
	for _, w := range wordsInBoard {
		words = append(words, w.Word)
	}
	sort.Slice(words, func(a, b int) bool { return words[a] < words[b] })
	expectedWords := []string{"hell", "hello"}
	if diff := deep.Equal(words, expectedWords); diff != nil {
		t.Error(diff)
	}
}

func BenchmarkWordsInBoard(b *testing.B) {
	var fh *os.File
	path := "contrib/opentaal-wordlist/wordlist.txt"
	for attempt := 0; attempt < 10; attempt++ {
		var err error
		fh, err = os.Open(path)
		if err == nil {
			break
		}
		path = "../" + path
	}
	if fh == nil {
		b.Skip("Cannot find or open opentaal-wordlist, skipping this benchmark")
	}

	board := vierkantle.NewBoard(5, 5)
	board.Cells[0] = []rune{'d', 'l', 'e', 'm', 'o'}
	board.Cells[1] = []rune{'a', 'l', 'o', 'd', 'a'}
	board.Cells[2] = []rune{'c', 'g', 'e', 'r', 'd'}
	board.Cells[3] = []rune{'m', 'y', 'i', 'n', 's'}
	board.Cells[4] = []rune{'i', 'i', 't', 'm', 'e'}

	for i := 0; i < b.N; i++ {
		dict := dictionary.NewPrefixDictionary()
		dict.Read(dictionary.NewFileReaderFromHandle(fh), dictionary.NormalWord, false)
		board.WordsInBoard(dict, 4)
	}
}
