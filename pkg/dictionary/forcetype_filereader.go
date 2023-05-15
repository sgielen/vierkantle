package dictionary

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type forceTypeFileReader struct {
	scanner *bufio.Scanner
}

func NewForceTypeFileReader(file string) (*forceTypeFileReader, error) {
	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return NewForceTypeFileReaderFromHandle(fh), nil
}

func NewForceTypeFileReaderFromHandle(fh *os.File) *forceTypeFileReader {
	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)
	return &forceTypeFileReader{scanner: scanner}
}

func (f *forceTypeFileReader) ReadWord() WordReadResult {
	if !f.scanner.Scan() {
		return WordReadResult{}
	}
	line := f.scanner.Text()
	words := strings.Fields(line)
	wordType := NoWord
	if len(words) > 1 {
		if words[1] == "no" {
			wordType = NoWord
		} else if words[1] == "normal" {
			wordType = NormalWord
		} else if words[1] == "bonus" {
			wordType = BonusWord
		} else if words[1] == "sensitive" {
			wordType = SensitiveWord
		}
	}
	return WordReadResult{
		Word:      words[0],
		ForceType: &wordType,
	}
}

func AddForceTypeToFile(file string, word string, wordType WordType) error {
	fh, err := os.Open(file)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fh, err = os.Create(file)
		}
		if err != nil {
			return err
		}
	}

	words := map[string]WordType{}
	reader := NewForceTypeFileReaderFromHandle(fh)
	for {
		word := reader.ReadWord()
		if word.Word == "" {
			break
		}
		words[word.Word] = *word.ForceType
	}
	words[word] = wordType
	if err := fh.Close(); err != nil {
		return err
	}

	fh, err = os.Create(file + ".bak")
	if err != nil {
		return err
	}

	for word, wordType := range words {
		var wordTypeStr string
		switch wordType {
		case NoWord:
			wordTypeStr = "no"
		case NormalWord:
			wordTypeStr = "normal"
		case BonusWord:
			wordTypeStr = "bonus"
		case SensitiveWord:
			wordTypeStr = "sensitive"
		}
		if _, err := fh.WriteString(fmt.Sprintf("%s %s\n", word, wordTypeStr)); err != nil {
			return err
		}
	}
	if err := fh.Close(); err != nil {
		return err
	}
	return os.Rename(file+".bak", file)
}
