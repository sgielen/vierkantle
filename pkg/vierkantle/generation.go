package vierkantle

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/sgielen/vierkantle/pkg/dictionary"
)

func (b *Board) randomPathWithLength(length int, path []Coord) []Coord {
	nextPaths := b.NextNavigationsFrom(path)

	random := rand.New(rand.NewSource(time.Now().UnixMilli()))
	random.Shuffle(len(nextPaths), func(i, j int) {
		nextPaths[i], nextPaths[j] = nextPaths[j], nextPaths[i]
	})

	for _, nextPath := range nextPaths {
		if len(nextPath) == length {
			return nextPath
		}
		if fullPath := b.randomPathWithLength(length, nextPath); fullPath != nil {
			return fullPath
		}
	}

	return nil
}

func (b *Board) Fill(c rune) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			b.Cells[y][x] = c
		}
	}
}

func (b *Board) PrefillRandomly(word string) error {
	path := b.randomPathWithLength(len(word), nil)
	if path == nil {
		return fmt.Errorf("the prefill word doesn't fit in the board :-(")
	}
	for i, c := range path {
		b.Cells[c.Y][c.X] = rune(word[i])
	}
	return nil
}

func (b *Board) FillRandomly() {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	random := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Cells[y][x] == '?' || b.Cells[y][x] == 0 {
				b.Cells[y][x] = letters[random.Intn(len(letters))]
			}
		}
	}
}

func (b *Board) ResetCell(coord Coord) {
	b.Cells[coord.Y][coord.X] = '?'
}

// Try to get this board filled up with random letters until all of them
// are used in words
func (b *Board) FillFullyUsed(dict dictionary.PrefixDictionary) ([]WordInBoard, bool) {
	subAttempts := 100
	var words []WordInBoard
	var unusedCells []Coord
	for subAttempt := 0; subAttempt < subAttempts; subAttempt++ {
		b.FillRandomly()
		words = b.WordsInBoard(dict, 4)
		unusedCells = b.FindUnusedCells(words)
		if len(unusedCells) == 0 {
			break
		}

		for _, unusedCell := range unusedCells {
			b.ResetCell(unusedCell)
		}
	}

	return words, len(unusedCells) == 0
}

func (b *Board) ScoreBoard(words []WordInBoard) float64 {
	score := 0.
	numNormalWords := 0
	for _, wordInBoard := range words {
		if wordInBoard.WordType == dictionary.NormalWord {
			numNormalWords++
			frequencyScore := math.Log10(wordInBoard.Frequency)
			if frequencyScore < 0 {
				frequencyScore = 0
			}
			lengthScore := float64(len(wordInBoard.Word)-4) * 3
			score += lengthScore * frequencyScore
		}
	}
	score = score / float64(numNormalWords)
	// Score is now the average word length throughout the board - longer is usually better.
	// Give a small penalty for having too many or too few words.
	if numNormalWords > 60 {
		score -= float64(numNormalWords-60) * 0.1
	} else if numNormalWords < 25 {
		score -= float64(25-numNormalWords) * 0.1
	} else {
		// more words is slightly better
		score += float64(numNormalWords) * 0.01
	}
	return score
}
