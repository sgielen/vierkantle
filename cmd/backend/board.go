package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sgielen/vierkantle/pkg/dictionary"
	pb "github.com/sgielen/vierkantle/pkg/proto"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var NoBoardsPresent = status.Error(codes.NotFound, "no boards present")

func (s *vierkantleService) GetBoard(ctx context.Context, req *pb.GetBoardRequest) (*pb.GetBoardResponse, error) {
	files, err := os.ReadDir(s.boardDir)
	if err != nil {
		return nil, NoBoardsPresent
	}

	// ReadDir returns sorted files, so we simply return the last .json file
	for i := len(files) - 1; i >= 0; i-- {
		if !strings.HasSuffix(files[i].Name(), ".json") {
			continue
		}
		bytes, err := os.ReadFile(filepath.Join(s.boardDir, files[i].Name()))
		if err != nil {
			s.log.Printf("failed to read board %s, skipping: %v", files[i].Name(), err)
			continue
		}
		return &pb.GetBoardResponse{
			Board: bytes,
		}, nil
	}

	return nil, NoBoardsPresent
}

func (s *vierkantleService) getDictionary() (dictionary.RWPrefixDictionary, error) {
	dict := dictionary.NewPrefixDictionary()
	for _, bonuslist := range s.bonusLists {
		if err := dict.ReadFromFile(bonuslist, dictionary.BonusWord, false); err != nil {
			return nil, err
		}
	}
	for _, wordlist := range s.wordLists {
		if err := dict.ReadFromFile(wordlist, dictionary.NormalWord, true /* upgrade only */); err != nil {
			return nil, err
		}
	}
	return dict, nil
}

func (s *vierkantleService) WordsForBoard(ctx context.Context, req *pb.WordsForBoardRequest) (*pb.WordsForBoardResponse, error) {
	dict, err := s.getDictionary()
	if err != nil {
		return nil, err
	}

	board, _, err := vierkantle.BoardFromJson(req.Board)
	if err != nil {
		return nil, err
	}
	words := board.WordsInBoard(dict, 4)
	newBoard, err := board.PrintBoardJson(words)
	if err != nil {
		return nil, err
	}
	return &pb.WordsForBoardResponse{
		Board: newBoard,
	}, nil
}

func (s *vierkantleService) SeedBoard(ctx context.Context, req *pb.SeedBoardRequest) (*pb.SeedBoardResponse, error) {
	dict, err := s.getDictionary()
	if err != nil {
		return nil, err
	}

	if req.SeedWord != "" {
		dict.AddWord(dictionary.WordReadResult{
			Word: req.SeedWord,
		}, dictionary.NormalWord, false)
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0.
	attempts := 5000
	for attempt := 0; attempt < attempts; attempt++ {
		board := vierkantle.NewBoard(int(req.Width), int(req.Height))
		if err := board.PrefillRandomly(req.SeedWord); err != nil {
			return nil, fmt.Errorf("that seed word doesn't fit in the board :-(")
		}

		words, ok := board.FillFullyUsed(dict)
		if !ok && attempt != attempts-1 /* last attempt */ {
			// Couldn't fill up this board to have words, nevermind
			continue
		}

		score := board.ScoreBoard(words)
		if score > bestScore {
			bestBoard = board
			bestWords = words
			bestScore = score
		}
	}

	newBoard, err := bestBoard.PrintBoardJson(bestWords)
	if err != nil {
		return nil, err
	}
	return &pb.SeedBoardResponse{
		Board: newBoard,
	}, nil
}

func (s *vierkantleService) FillInBoard(ctx context.Context, req *pb.FillInBoardRequest) (*pb.FillInBoardResponse, error) {
	dict, err := s.getDictionary()
	if err != nil {
		return nil, err
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0.
	attempts := 5000
	for attempt := 0; attempt < attempts; attempt++ {
		board, _, err := vierkantle.BoardFromJson(req.Board)
		if err != nil {
			return nil, err
		}

		words, ok := board.FillFullyUsed(dict)
		if !ok && attempt != attempts-1 /* last attempt */ {
			// Couldn't fill up this board to have words, nevermind
			continue
		}

		score := board.ScoreBoard(words)
		if score > bestScore {
			bestBoard = board
			bestWords = words
			bestScore = score
		}
	}
	newBoard, err := bestBoard.PrintBoardJson(bestWords)
	if err != nil {
		return nil, err
	}
	return &pb.FillInBoardResponse{
		Board: newBoard,
	}, nil
}
