package vierkantle

import "fmt"

type Board struct {
	Width  int
	Height int
	Cells  [][]rune
}

func NewBoard(width, height int) *Board {
	cells := make([][]rune, width)
	for x := range cells {
		cells[x] = make([]rune, height)
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
			res = res + fmt.Sprintf("%c   ", b.Cells[x][y])
		}
		res = res + "  \n"
	}
	return res + "\n"
}

func (b *Board) PrintBoardGo() string {
	res := "\n"
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			res = res + fmt.Sprintf("board.Cells[%d][%d] = '%c'\n", x, y, b.Cells[x][y])
		}
	}
	return res + "\n"
}

func (b *Board) PrintBoardJs() string {
	res := "\n{\n"
	res = res + fmt.Sprintf("  width: %d,\n", b.Width)
	res = res + fmt.Sprintf("  height: %d,\n", b.Height)
	res = res + "  cells: [\n"
	for y := 0; y < b.Height; y++ {
		res = res + "    ["
		for x := 0; x < b.Width; x++ {
			res = res + fmt.Sprintf("'%c', ", b.Cells[x][y])
		}
		res = res + "],\n"
	}
	res = res + "  ],\n"
	res = res + "}\n"

	return res
}
