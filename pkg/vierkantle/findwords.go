package vierkantle

import (
	"github.com/sgielen/vierkantle/pkg/dictionary"
)

type WordInBoard struct {
	Word     string
	Path     Path
	WordType dictionary.WordType
}

func (b *Board) WordForPath(path Path) string {
	res := make([]rune, len(path))
	for i, c := range path {
		res[i] = b.Cells[c.X][c.Y]
	}
	return string(res)
}

func (b *Board) WordsInBoard(dictionary dictionary.PrefixDictionary, minLength int) []WordInBoard {
	words := map[string]WordInBoard{}
	b.wordsInBoard(&words, dictionary, minLength, nil)
	res := make([]WordInBoard, len(words))[:0]
	for _, wordInBoard := range words {
		res = append(res, wordInBoard)
	}
	return res
}

func (b *Board) wordsInBoard(words *map[string]WordInBoard, dict dictionary.PrefixDictionary, minLength int, path Path) {
	nextPaths := b.NextNavigationsFrom(path)
	for _, nextPath := range nextPaths {
		word := b.WordForPath(nextPath)
		wordMatch := dict.HasWord(word)
		if wordMatch.HasThisWord != dictionary.NoWord && len(nextPath) >= minLength {
			(*words)[word] = WordInBoard{
				Word:     word,
				Path:     nextPath,
				WordType: wordMatch.HasThisWord,
			}
		}
		if wordMatch.HasWordsWithPrefix {
			b.wordsInBoard(words, dict, minLength, nextPath)
		}
	}
}

func (b *Board) AreAllCellsUsed(words []WordInBoard) bool {
	cells := make([][]bool, b.Width)
	for x := range cells {
		cells[x] = make([]bool, b.Height)
		for y := range cells[x] {
			cells[x][y] = false
		}
	}
	cellsRemaining := b.Width * b.Height
	for _, word := range words {
		if word.WordType != dictionary.NormalWord {
			continue
		}
		for _, coord := range word.Path {
			if cells[coord.X][coord.Y] == false {
				cells[coord.X][coord.Y] = true
				cellsRemaining -= 1
				if cellsRemaining == 0 {
					return true
				}
			}
		}
	}
	return false
}
