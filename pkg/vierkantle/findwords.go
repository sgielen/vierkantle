package vierkantle

import (
	"github.com/sgielen/vierkantle/pkg/dictionary"
)

func (b *Board) WordForPath(path Path) string {
	res := make([]rune, len(path))
	for i, c := range path {
		res[i] = b.Cells[c.X][c.Y]
	}
	return string(res)
}

func (b *Board) WordsInBoard(dictionary dictionary.PrefixDictionary, minLength int) []string {
	words := map[string]struct{}{}
	b.wordsInBoard(&words, dictionary, minLength, nil)
	res := make([]string, len(words))[:0]
	for w := range words {
		res = append(res, w)
	}
	return res
}

func (b *Board) wordsInBoard(words *map[string]struct{}, dictionary dictionary.PrefixDictionary, minLength int, path Path) {
	nextPaths := b.NextNavigationsFrom(path)
	for _, nextPath := range nextPaths {
		word := b.WordForPath(nextPath)
		wordMatch := dictionary.HasWord(word)
		if wordMatch.HasThisWord && len(nextPath) >= minLength {
			(*words)[word] = struct{}{}
		}
		if wordMatch.HasWordsWithPrefix {
			b.wordsInBoard(words, dictionary, minLength, nextPath)
		}
	}
}
