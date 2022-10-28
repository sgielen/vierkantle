package vierkantle_test

import (
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
			res.HasThisWord = true
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

	words := board.WordsInBoard(mockDictionary, 4)
	expectedWords := []string{"hell", "hello"}
	if diff := deep.Equal(words, expectedWords); diff != nil {
		t.Error(diff)
	}
}
