package vierkantle

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixMilli()))
}

func (b *Board) randomPathWithLength(length int, path []Coord) []Coord {
	nextPaths := b.NextNavigationsFrom(path)

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
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			b.Cells[x][y] = c
		}
	}
}

func (b *Board) PrefillRandomly(word string) {
	path := b.randomPathWithLength(len(word), nil)
	for i, c := range path {
		b.Cells[c.X][c.Y] = rune(word[i])
	}
}

func (b *Board) FillRandomly() {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			if b.Cells[x][y] == '?' {
				b.Cells[x][y] = letters[random.Intn(len(letters))]
			}
		}
	}
}
