package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sgielen/vierkantle/pkg/vierkantle"
)

func boardDir() string {
	if dir := os.Getenv("BOARDS_DIR"); dir != "" {
		return dir
	} else {
		return "."
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %s board", os.Args[0])
	}

	filename := os.Args[1]
	abs, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalf("Error abs filename: %v", err)
	}

	basename := filepath.Base(filename)
	basename, _ = strings.CutSuffix(basename, ".json")

	board, words, err := vierkantle.BoardFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading board: %v", err)
	}

	log.Printf("BOARD: %s\n", filename)
	log.Printf("Board name: %s\n", basename)
	log.Printf("Longest words: %v\n", board.LongestWords(words, 4))
	letters := board.PrintBoard()
	log.Print(letters)
	log.Printf("Score: %f\n", board.ScoreBoard(words))

	var instances []string
	if err := filepath.WalkDir(boardDir(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		a, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		if a == abs {
			return nil
		}
		b, _, err := vierkantle.BoardFromFile(path)
		if err == nil && b.PrintBoard() == letters {
			instances = append(instances, filepath.Base(path))
		}
		return err
	}); err != nil {
		log.Fatalf("Error finding other uses of board: %v", err)
	}
	log.Printf("Other instances of this board: %v", instances)
}
