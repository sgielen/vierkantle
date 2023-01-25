package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sgielen/vierkantle/pkg/dictionary"
	pb "github.com/sgielen/vierkantle/pkg/proto"
	"github.com/sgielen/vierkantle/pkg/vierkantle"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var NoBoardsPresent = status.Error(codes.NotFound, "no boards present")

func (s *vierkantleService) GetBoard(ctx context.Context, req *pb.GetBoardRequest) (*pb.GetBoardResponse, error) {
	if req.TimezoneOffsetMinutes < -24*60 || req.TimezoneOffsetMinutes > 24*60 {
		return nil, status.Error(codes.InvalidArgument, "invalid timezone offset")
	}
	date := time.Now().UTC().In(time.FixedZone("", int(60*-req.TimezoneOffsetMinutes)))
	// If it's before 12:00 PM in the user's timezone, take yesterday's board
	if date.Hour() < 12 {
		date = date.Add(-time.Hour * 13)
	}
	todaysBoard := date.Format("2006-01-02") + ".json"
	bytes, err := os.ReadFile(filepath.Join(s.boardDir, todaysBoard))
	if err == nil {
		return &pb.GetBoardResponse{
			Board: bytes,
		}, nil
	}

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

func (s *vierkantleService) SeedBoard(req *pb.SeedBoardRequest, stream pb.VierkantleService_SeedBoardServer) error {
	dict, err := s.getDictionary()
	if err != nil {
		return err
	}

	if req.SeedWord != "" {
		dict.AddWord(dictionary.WordReadResult{
			Word: req.SeedWord,
		}, dictionary.NormalWord, false)
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0.
	sendInterval := req.Attempts / 100
	var nextSend int32
	for attempt := int32(0); attempt < req.Attempts; attempt++ {
		board := vierkantle.NewBoard(int(req.Width), int(req.Height))
		if err := board.PrefillRandomly(req.SeedWord); err != nil {
			return fmt.Errorf("that seed word doesn't fit in the board :-(")
		}

		words, ok := board.FillFullyUsed(dict)
		if !ok && attempt != req.Attempts-1 /* last attempt */ {
			// Couldn't fill up this board to have words, nevermind
			continue
		}

		score := board.ScoreBoard(words)
		if score > bestScore {
			bestBoard = board
			bestWords = words
			bestScore = score
		}

		if nextSend <= 0 && bestBoard != nil {
			nextSend = sendInterval
			newBoard, err := bestBoard.PrintBoardJson(bestWords)
			if err != nil {
				return err
			}
			if err := stream.Send(&pb.SeedBoardResponse{
				Board:    newBoard,
				Attempts: req.Attempts,
				Progress: attempt,
			}); err != nil {
				return err
			}
		} else {
			nextSend -= 1
		}
	}

	newBoard, err := bestBoard.PrintBoardJson(bestWords)
	if err != nil {
		return err
	}
	return stream.Send(&pb.SeedBoardResponse{
		Board:    newBoard,
		Attempts: req.Attempts,
		Progress: req.Attempts,
	})
}

func (s *vierkantleService) FillInBoard(req *pb.FillInBoardRequest, stream pb.VierkantleService_FillInBoardServer) error {
	dict, err := s.getDictionary()
	if err != nil {
		return err
	}

	var bestBoard *vierkantle.Board
	var bestWords []vierkantle.WordInBoard
	bestScore := 0.
	sendInterval := req.Attempts / 100
	var nextSend int32
	for attempt := int32(0); attempt < req.Attempts; attempt++ {
		board, _, err := vierkantle.BoardFromJson(req.Board)
		if err != nil {
			return err
		}

		words, ok := board.FillFullyUsed(dict)
		if !ok && attempt != req.Attempts-1 /* last attempt */ {
			// Couldn't fill up this board to have words, nevermind
			continue
		}

		score := board.ScoreBoard(words)
		if score > bestScore {
			bestBoard = board
			bestWords = words
			bestScore = score
		}

		if nextSend <= 0 && bestBoard != nil {
			nextSend = sendInterval
			newBoard, err := bestBoard.PrintBoardJson(bestWords)
			if err != nil {
				return err
			}
			if err := stream.Send(&pb.FillInBoardResponse{
				Board:    newBoard,
				Attempts: req.Attempts,
				Progress: attempt,
			}); err != nil {
				return err
			}
		} else {
			nextSend -= 1
		}
	}

	newBoard, err := bestBoard.PrintBoardJson(bestWords)
	if err != nil {
		return err
	}
	return stream.Send(&pb.FillInBoardResponse{
		Board:    newBoard,
		Attempts: req.Attempts,
		Progress: req.Attempts,
	})
}
