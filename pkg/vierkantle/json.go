package vierkantle

import (
	"encoding/json"

	"github.com/sgielen/vierkantle/pkg/dictionary"
)

type BoardJsonWord struct {
	Path  Path `json:"path"`
	Bonus bool `json:"bonus,omitempty"`
	Swear bool `json:"swear,omitempty"`
}

type BoardJsonExport struct {
	Width  int                      `json:"width"`
	Height int                      `json:"height"`
	Cells  [][]string               `json:"cells"` // [y][x]
	Words  map[string]BoardJsonWord `json:"words"`
}

func (b *Board) PrintBoardJson(words []WordInBoard) ([]byte, error) {
	res := BoardJsonExport{}
	res.Width = b.Width
	res.Height = b.Height
	res.Cells = make([][]string, res.Width)
	for y, row := range b.Cells {
		res.Cells[y] = make([]string, res.Height)
		for x, cell := range row {
			res.Cells[y][x] = string(cell)
		}
	}
	res.Words = make(map[string]BoardJsonWord)
	for _, word := range words {
		res.Words[word.Word] = BoardJsonWord{
			Path:  word.Path,
			Bonus: word.WordType == dictionary.BonusWord,
			Swear: word.WordType == dictionary.SwearWord,
		}
	}
	return json.Marshal(res)
}
