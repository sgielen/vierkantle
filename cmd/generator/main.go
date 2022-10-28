package main

import (
	"log"
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
	var bestWords []string
	bestScore := 0

	niceWord := "algoritmes"
	attempts := 100
	for attempt := 0; attempt < attempts; attempt++ {
		board := vierkantle.NewBoard(boardWidth, boardHeight)
		board.PrefillRandomly(niceWord)
		board.FillRandomly()

		words := board.WordsInBoard(dictionary, 4)
		score := 0
		for _, word := range words {
			score += len(word) * len(word)
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
	//log.Print("\n" + board.PrintBoardGo())

	log.Printf("Analyzing board...")
	wordsByLength := map[int][]string{}
	for _, word := range words {
		wordsByLength[len(word)] = append(wordsByLength[len(word)], word)
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
