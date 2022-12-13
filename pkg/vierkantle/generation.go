package vierkantle

import (
	"fmt"
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
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Cells[y][x] == '?' {
				b.Cells[y][x] = letters[random.Intn(len(letters))]
			}
		}
	}
}

func (b *Board) ResetCell(coord Coord) {
	b.Cells[coord.Y][coord.X] = '?'
}
