package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/sgielen/vierkantle/pkg/dictionary"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
)

func readUntilArg(args []string, list *[]string) int {
	for i := 0; i < len(args); i += 1 {
		if len(args[i]) > 0 && args[i][0] == '-' {
			return i
		}
		*list = append(*list, args[i])
	}
	return len(args)
}

func help() {
	log.Fatalf("Usage: %s -s 4 -w seedword -W wordlist [wordlist..] -B bonuslist [bonuslist..] -o board.json", os.Args[0])
}

func main() {
	var wordlists []string
	var bonuslists []string
	output := "board.json"
	var seedWord string
	size := 4

	for i := 1; i < len(os.Args); {
		switch os.Args[i] {
		case "-W":
			i += 1
			i += readUntilArg(os.Args[i:], &wordlists)
		case "-B":
			i += 1
			i += readUntilArg(os.Args[i:], &bonuslists)
		case "-w":
			seedWord = os.Args[i+1]
			i += 2
		case "-o":
			output = os.Args[i+1]
			i += 2
		case "-s":
			var err error
			size, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				log.Print(err)
				help()
			}
			i += 2
		default:
			log.Printf("Unknown parameter: %s", os.Args[i])
			fallthrough
		case "-h":
			help()
		}
	}

	if len(wordlists) == 0 {
		log.Printf("No wordlists given.")
		help()
	}

	boardWidth := size
	boardHeight := size

	log.Printf("Reading dictionaries...")
	dict := dictionary.NewPrefixDictionary()
	for _, bonuslist := range bonuslists {
		if err := dict.ReadFromFile(bonuslist, dictionary.BonusWord, false); err != nil {
			log.Fatal(err)
		}
	}
	for _, wordlist := range wordlists {
		if err := dict.ReadFromFile(wordlist, dictionary.NormalWord, true /* upgrade only */); err != nil {
			log.Fatal(err)
		}
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0.

	log.Printf("Generating boards...")
	if seedWord != "" {
		dict.AddWord(dictionary.WordReadResult{
			Word: seedWord,
		}, dictionary.NormalWord, false)
	}
	attempts := 10000
	for attempt := 0; attempt < attempts; attempt++ {
		board := vierkantle.NewBoard(boardWidth, boardHeight)
		if err := board.PrefillRandomly(seedWord); err != nil {
			log.Fatal("That seed word doesn't fit in the board :-(")
		}

		// Try to get this board filled up with random letters
		// until it has words
		subAttempts := 100
		var words []vierkantle.WordInBoard
		var unusedCells []vierkantle.Coord
		for subAttempt := 0; subAttempt < subAttempts; subAttempt++ {
			board.FillRandomly()
			words = board.WordsInBoard(dict, 4)
			unusedCells = board.FindUnusedCells(words)
			if len(unusedCells) == 0 {
				break
			}

			for _, unusedCell := range unusedCells {
				board.ResetCell(unusedCell)
			}
		}

		if len(unusedCells) != 0 {
			// Couldn't fill up this board to have words, nevermind
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

	if fh, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
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
