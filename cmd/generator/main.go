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
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s bonuslist wordlist", os.Args[0])
	}
	bonuslist := os.Args[1]
	wordlist := os.Args[2]

	boardWidth := 4
	boardHeight := 4

	log.Printf("Reading dictionaries...")
	dict := dictionary.NewPrefixDictionary()
	if err := dict.ReadFromFile(bonuslist, dictionary.BonusWord, false); err != nil {
		log.Fatal(err)
	}
	if err := dict.ReadFromFile(wordlist, dictionary.NormalWord, true /* upgrade only */); err != nil {
		log.Fatal(err)
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0.

	log.Printf("Generating boards...")
	niceWord := "algoritmes"
	attempts := 10000
	for attempt := 0; attempt < attempts; attempt++ {
		board := vierkantle.NewBoard(boardWidth, boardHeight)
		if err := board.PrefillRandomly(niceWord); err != nil {
			log.Fatal("That prefill word doesn't fit in the board :-(")
		}
		board.FillRandomly()

		words := board.WordsInBoard(dict, 4)
		if !board.AreAllCellsUsed(words) {
			// nevermind, skip this board
			continue
		}

		score := 0.
		numNormalWords := 0
		for _, wordInBoard := range words {
			if wordInBoard.WordType == dictionary.NormalWord {
				numNormalWords++
				frequencyScore := math.Log10(wordInBoard.Frequency)
				if frequencyScore < 0 {
					frequencyScore = 0
				}
				lengthScore := float64(len(wordInBoard.Word)-4) * 3
				score += lengthScore * frequencyScore
			}
		}
		score = score / float64(numNormalWords)
		if score > bestScore {
			bestBoard = board
			bestWords = words
			bestScore = score
		}
	}

	if bestBoard == nil {
		log.Fatal("failed to generate a board")
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
	wordsByLength := map[int][]vierkantle.WordInBoard{}
	for _, word := range words {
		wordsByLength[len(word.Word)] = append(wordsByLength[len(word.Word)], word)
	}
	var lengths []int
	for len := range wordsByLength {
		lengths = append(lengths, len)
	}
	sort.Ints(lengths)
	letters := 0
	normalWords := 0
	for _, length := range lengths {
		words := wordsByLength[length]
		log.Printf("%d words of length %d:", len(words), length)
		sort.Slice(words, func(a, b int) bool { return words[a].Word < words[b].Word })
		for _, w := range words {
			tp := ""
			if w.WordType == dictionary.BonusWord {
				tp = " (bonus)"
			} else if w.WordType == dictionary.SwearWord {
				tp = " (swear)"
			}
			log.Printf("- %s%s", w.Word, tp)
			if w.WordType == dictionary.NormalWord {
				normalWords += 1
				letters += len(w.Word)
			}
		}
	}
	log.Printf("%d normal words in total, %d letters in total!", normalWords, letters)
}
