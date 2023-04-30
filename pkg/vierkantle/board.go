package vierkantle

import "fmt"

type Board struct {
	Width  int
	Height int
	Cells  [][]rune // [y][x]
}

func NewBoard(width, height int) *Board {
	cells := make([][]rune, height)
	for y := range cells {
		cells[y] = make([]rune, width)
	}
	board := &Board{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
	board.Fill('?')
	return board
}

func (b *Board) PrintBoard() string {
	res := "\n"
	for y := 0; y < b.Height; y++ {
		res = res + "  "
		for x := 0; x < b.Width; x++ {
			res = res + fmt.Sprintf("%c   ", b.Cells[y][x])
		}
		res = res + "  \n"
	}
	return res + "\n"
}

func (b *Board) PrintBoardGo() string {
	res := "\n"
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			res = res + fmt.Sprintf("board.Cells[%d][%d] = '%c'\n", y, x, b.Cells[y][x])
		}
	}
	return res + "\n"
}
