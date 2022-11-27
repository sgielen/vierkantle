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
	Cells  [][]string               `json:"cells"`
	Words  map[string]BoardJsonWord `json:"words"`
}

func (b *Board) PrintBoardJson(words []WordInBoard) ([]byte, error) {
	res := BoardJsonExport{}
	res.Width = b.Width
	res.Height = b.Height
	res.Cells = make([][]string, res.Width)
	for x, row := range b.Cells {
		res.Cells[x] = make([]string, res.Height)
		for y, cell := range row {
			res.Cells[x][y] = string(cell)
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
