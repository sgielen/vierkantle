package vierkantle

import (
	"github.com/sgielen/vierkantle/pkg/dictionary"
)

type WordInBoard struct {
	Word      string
	Path      Path
	WordType  dictionary.WordType
	Frequency float64
}

func (b *Board) WordForPath(path Path) string {
	res := make([]rune, len(path))
	for i, c := range path {
		res[i] = b.Cells[c.Y][c.X]
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
				Word:      word,
				Path:      nextPath,
				WordType:  wordMatch.HasThisWord,
				Frequency: wordMatch.Frequency,
			}
		}
		if wordMatch.HasWordsWithPrefix {
			b.wordsInBoard(words, dict, minLength, nextPath)
		}
	}
}

func (b *Board) AreAllCellsUsed(words []WordInBoard) bool {
	cells := make([][]bool, b.Height)
	for y := range cells {
		cells[y] = make([]bool, b.Width)
		for x := range cells[y] {
			cells[y][x] = false
		}
	}
	cellsRemaining := b.Width * b.Height
	for _, word := range words {
		if word.WordType != dictionary.NormalWord {
			continue
		}
		for _, coord := range word.Path {
			if !cells[coord.Y][coord.X] {
				cells[coord.Y][coord.X] = true
				cellsRemaining -= 1
				if cellsRemaining == 0 {
					return true
				}
			}
		}
	}
	return false
}
