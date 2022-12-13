package dictionary

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type fileReader struct {
	scanner *bufio.Scanner
}

func NewFileReader(file string) (*fileReader, error) {
	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return NewFileReaderFromHandle(fh), nil
}

func NewFileReaderFromHandle(fh *os.File) *fileReader {
	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)
	return &fileReader{scanner: scanner}
}

func (f *fileReader) ReadWord() WordReadResult {
	if !f.scanner.Scan() {
		return WordReadResult{}
	}
	line := f.scanner.Text()
	words := strings.Fields(line)
	var frequency float64
	if len(words) > 1 {
		frequency, _ = strconv.ParseFloat(words[1], 64)
	}
	return WordReadResult{
		Word:      words[0],
		Frequency: frequency,
	}
}
