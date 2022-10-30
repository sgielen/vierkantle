package main

import (
	"log"
	"math"
	"os"
	"sort"

	"github.com/sgielen/vierkantle/pkg/dictionary"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s wordlist", os.Args[0])
	}
	wordlist := os.Args[1]

	boardWidth := 4
	boardHeight := 4

	log.Printf("Reading dictionary...")
	dictionary, err := dictionary.NewPrefixDictionaryFromFile(wordlist, 4, boardWidth*boardHeight)
	if err != nil {
		log.Fatal(err)
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0

	log.Printf("Generating boards...")
	niceWord := "algoritmes"
	attempts := 10000
	for attempt := 0; attempt < attempts; attempt++ {
		board := vierkantle.NewBoard(boardWidth, boardHeight)
		if err := board.PrefillRandomly(niceWord); err != nil {
			log.Fatal("That prefill word doesn't fit in the board :-(")
		}
		board.FillRandomly()

		words := board.WordsInBoard(dictionary, 4)
		if !board.AreAllCellsUsed(words) {
			// nevermind, skip this board
			continue
		}
		score := 0
		for _, wordInBoard := range words {
			score += int(math.Pow(float64(len(wordInBoard.Word)), 3))
		}
		if score > bestScore {
			bestBoard = board
			bestWords = words
			bestScore = score
		}
	}

	board := bestBoard
	words := bestWords

	log.Print("\n" + board.PrintBoard())
	log.Print("\n" + board.PrintBoardGo())

	json, err := board.PrintBoardJson(words)
	if err != nil {
		log.Fatalf("failed to generate JSON for board: %s", err.Error())
	}

	if fh, err := os.OpenFile("board.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
		log.Fatalf("failed to open file to write JSON: %s", err.Error())
	} else if _, err = fh.Write(json); err != nil {
		log.Fatalf("failed to write JSON to file: %s", err.Error())
	} else if err = fh.Close(); err != nil {
		log.Fatalf("failed to close file: %s", err.Error())
	}

	log.Printf("Analyzing board...")
	wordsByLength := map[int][]string{}
	for _, word := range words {
		wordsByLength[len(word.Word)] = append(wordsByLength[len(word.Word)], word.Word)
	}
	var lengths []int
	for len := range wordsByLength {
		lengths = append(lengths, len)
	}
	sort.Ints(lengths)
	letters := 0
	for _, length := range lengths {
		words := wordsByLength[length]
		log.Printf("%d words of length %d:", len(words), length)
		sort.Strings(words)
		for _, w := range words {
			log.Printf("- %s", w)
			letters += len(w)
		}
	}
	log.Printf("%d words in total, %d letters in total!", len(words), letters)
}
