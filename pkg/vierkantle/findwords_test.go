package vierkantle_test

import (
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/sgielen/vierkantle/pkg/dictionary"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
)

func TestWordForPath(t *testing.T) {
	board := vierkantle.NewBoard(3, 3)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			board.Cells[x][y] = 'x'
		}
	}
	board.Cells[0][0] = 'h'
	board.Cells[0][1] = 'e'
	board.Cells[1][2] = 'l'
	board.Cells[2][2] = 'l'
	board.Cells[2][1] = 'o'

	word := board.WordForPath(vierkantle.Path{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}})
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
	board.Cells[0][0] = 'h'
	board.Cells[0][1] = 'e'
	board.Cells[0][2] = 'x'
	board.Cells[1][0] = /**/ 'x'
	board.Cells[1][1] = /**/ 'i'
	board.Cells[1][2] = /**/ 'l'
	board.Cells[2][0] = /**/ /**/ 'x'
	board.Cells[2][1] = /**/ /**/ 'o'
	board.Cells[2][2] = /**/ /**/ 'l'

	mockDictionary := &MockDictionary{}
	mockDictionary.ExistingWords = []string{"hell", "hello", "hellish"}
	mockDictionary.ExistingPrefixes = []string{"h", "he", "hel", "hell", "hello", "helli", "hellis", "hellish"}

	wordsInBoard := board.WordsInBoard(mockDictionary, 4)
	var words []string
	for _, w := range wordsInBoard {
		words = append(words, w.Word)
	}
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
	board.Cells[0][0] = 'd'
	board.Cells[1][0] = 'l'
	board.Cells[2][0] = 'e'
	board.Cells[3][0] = 'm'
	board.Cells[4][0] = 'o'
	board.Cells[0][1] = 'a'
	board.Cells[1][1] = 'l'
	board.Cells[2][1] = 'o'
	board.Cells[3][1] = 'd'
	board.Cells[4][1] = 'a'
	board.Cells[0][2] = 'c'
	board.Cells[1][2] = 'g'
	board.Cells[2][2] = 'e'
	board.Cells[3][2] = 'r'
	board.Cells[4][2] = 'd'
	board.Cells[0][3] = 'm'
	board.Cells[1][3] = 'y'
	board.Cells[2][3] = 'i'
	board.Cells[3][3] = 'n'
	board.Cells[4][3] = 's'
	board.Cells[0][4] = 'i'
	board.Cells[1][4] = 'i'
	board.Cells[2][4] = 't'
	board.Cells[3][4] = 'm'
	board.Cells[4][4] = 'e'

	for i := 0; i < b.N; i++ {
		dict := dictionary.NewPrefixDictionary()
		dict.Read(dictionary.NewFileReaderFromHandle(fh), dictionary.NormalWord, false)
		board.WordsInBoard(dict, 4)
	}
}
