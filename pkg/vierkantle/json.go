package vierkantle

import (
	"encoding/json"
	"os"

	"github.com/sgielen/vierkantle/pkg/dictionary"
)

type BoardJsonWord struct {
	Path      Path `json:"path"`
	Bonus     bool `json:"bonus,omitempty"`
	Sensitive bool `json:"sensitive,omitempty"`
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
	res.Cells = make([][]string, res.Height)
	for y, row := range b.Cells {
		res.Cells[y] = make([]string, res.Width)
		for x, cell := range row {
			if cell == 0 {
				res.Cells[y][x] = ""
			} else {
				res.Cells[y][x] = string(cell)
			}
		}
	}
	res.Words = make(map[string]BoardJsonWord)
	for _, word := range words {
		res.Words[word.Word] = BoardJsonWord{
			Path:      word.Path,
			Bonus:     word.WordType == dictionary.BonusWord,
			Sensitive: word.WordType == dictionary.SensitiveWord,
		}
	}
	return json.Marshal(res)
}

func BoardFromFile(filename string) (*Board, []WordInBoard, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	return BoardFromJson(data)
}

func BoardFromJson(data []byte) (*Board, []WordInBoard, error) {
	var req BoardJsonExport
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, nil, err
	}

	b := &Board{
		Width:  req.Width,
		Height: req.Height,
		Cells:  make([][]rune, req.Height),
	}
	for y, row := range req.Cells {
		b.Cells[y] = make([]rune, req.Width)
		for x, cell := range row {
			if len(cell) > 0 {
				b.Cells[y][x] = rune(cell[0])
			}
		}
	}
	var words []WordInBoard
	for word, wordInBoard := range req.Words {
		var wordType dictionary.WordType
		if wordInBoard.Bonus {
			wordType = dictionary.BonusWord
		} else if wordInBoard.Sensitive {
			wordType = dictionary.SensitiveWord
		} else {
			wordType = dictionary.NormalWord
		}
		words = append(words, WordInBoard{
			Word:      word,
			Path:      wordInBoard.Path,
			WordType:  wordType,
			Frequency: 0,
		})
	}
	return b, words, nil
}
