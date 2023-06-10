package main

import (
	"context"
	"fmt"
	"log"
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

func (s *vierkantleService) getNewestBoard() (string, []byte, error) {
	files, err := os.ReadDir(s.boardDir)
	if err != nil {
		return "", nil, NoBoardsPresent
	}

	// ReadDir returns sorted files, so we simply return the last .json file
	for i := len(files) - 1; i >= 0; i-- {
		name := files[i].Name()
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		bytes, err := os.ReadFile(filepath.Join(s.boardDir, name))
		if err != nil {
			s.log.Printf("failed to read board %s, skipping: %v", name, err)
			continue
		}
		return name[0 : len(name)-5], bytes, nil
	}

	return "", nil, NoBoardsPresent
}

func (s *vierkantleService) GetBoard(ctx context.Context, req *pb.GetBoardRequest) (*pb.GetBoardResponse, error) {
	if req.TimezoneOffsetMinutes < -24*60 || req.TimezoneOffsetMinutes > 24*60 {
		return nil, status.Error(codes.InvalidArgument, "invalid timezone offset")
	}
	date := time.Now().UTC().In(time.FixedZone("", int(60*-req.TimezoneOffsetMinutes)))
	// If it's before 12:00 PM in the user's timezone, take yesterday's board
	if date.Hour() < 12 {
		date = date.Add(-time.Hour * 13)
	}
	todaysBoard := date.Format("2006-01-02")
	bytes, err := os.ReadFile(filepath.Join(s.boardDir, todaysBoard+".json"))
	if err == nil {
		return &pb.GetBoardResponse{
			Board: bytes,
			Name:  todaysBoard,
		}, nil
	}

	newestBoard, bytes, err := s.getNewestBoard()
	return &pb.GetBoardResponse{
		Board: bytes,
		Name:  newestBoard,
	}, nil
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
	if s.forceWordTypeList != "" {
		if err := dict.ReadForceTypeFromFile(s.forceWordTypeList); err != nil {
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
	if req.Width*req.Height > 36 {
		return fmt.Errorf("refusing to SeedBoard for a board with >36 cells")
	}

	dict, err := s.getDictionary()
	if err != nil {
		return err
	}

	if req.SeedWord != "" {
		dict.AddWord(dictionary.WordReadResult{
			Word: req.SeedWord,
		}, dictionary.NormalWord, false, true)
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

	if bestBoard == nil {
		return fmt.Errorf("failed to generate any board")
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

		if board.Width*board.Height > 36 {
			return fmt.Errorf("refusing to FillInBoard for a board with >36 cells")
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

	if bestBoard == nil {
		return fmt.Errorf("failed to generate any board")
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

func (s *vierkantleService) MarkWordType(ctx context.Context, req *pb.MarkWordTypeRequest) (*pb.MarkWordTypeResponse, error) {
	// This RPC requires login, a slight measure against abuse
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	_, username, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("User %q marked word %q as bonus=%v", username, req.Word, req.Bonus)

	if s.forceWordTypeList == "" {
		return nil, fmt.Errorf("no force type word file set")
	}

	wordType := dictionary.NormalWord
	if req.Bonus {
		wordType = dictionary.BonusWord
	}

	if err := dictionary.AddForceTypeToFile(s.forceWordTypeList, req.Word, wordType); err != nil {
		return nil, err
	}
	return &pb.MarkWordTypeResponse{}, nil
}
