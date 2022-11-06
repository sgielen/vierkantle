package dictionary

import (
	"bufio"
	"os"
)

type fileReader struct {
	scanner *bufio.Scanner
}

func NewFileReader(file string) (WordReader, error) {
	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return NewFileReaderFromHandle(fh), nil
}

func NewFileReaderFromHandle(fh *os.File) WordReader {
	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)
	return &fileReader{scanner: scanner}
}

func (f *fileReader) ReadWord() string {
	if !f.scanner.Scan() {
		return ""
	}
	return f.scanner.Text()
}
