package main

import (
	"context"
	"fmt"
	"io/fs"
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

func (s *vierkantleService) GetTodaysBoard(ctx context.Context, date string) (*pb.GetBoardResponse, error) {
	bytes, err := os.ReadFile(s.GetBoardPath(date))
	if err != nil {
		return nil, err
	}
	return &pb.GetBoardResponse{
		Board: bytes,
		Name:  date,
	}, nil
}

func (s *vierkantleService) GetBoardPath(date string) string {
	return filepath.Join(s.boardDir, date+".json")
}

func (s *vierkantleService) GetEmergencyBoard() (string, error) {
	// letters -> latest filename
	boardMostRecentUse := map[string]string{}

	if err := filepath.WalkDir(s.boardDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		b, _, err := vierkantle.BoardFromFile(path)
		if err != nil {
			return err
		}
		letters := b.PrintBoard()
		if currentMostRecentUse, ok := boardMostRecentUse[letters]; ok {
			// Board was used before -- but earlier or later than this use?
			if path > currentMostRecentUse {
				boardMostRecentUse[letters] = path
			}
		} else {
			boardMostRecentUse[letters] = path
		}
		return err
	}); err != nil {
		return "", err
	}

	var oldestBoard string
	for _, path := range boardMostRecentUse {
		if oldestBoard == "" {
			oldestBoard = path
		} else if path < oldestBoard {
			oldestBoard = path
		}
	}

	if oldestBoard == "" {
		return "", fmt.Errorf("no boards to serve as emergency board")
	}

	oldestBoard = filepath.Base(oldestBoard)
	oldestBoard, _ = strings.CutSuffix(oldestBoard, ".json")
	return oldestBoard, nil
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
	today := date.Format("2006-01-02")
	res, err := s.GetTodaysBoard(ctx, today)
	if err == nil {
		return res, nil
	}

	// TODO: this mtx only works while the backend is single-process -- if the backend
	// becomes redundant, we should switch to a database lock
	s.emergencyBoardMtx.Lock()
	defer s.emergencyBoardMtx.Unlock()

	// Now that we have the lock, check again if there is a board for today
	res, err = s.GetTodaysBoard(ctx, today)
	if err == nil {
		return res, nil
	}

	// If not, we now have the lock to copy an emergency board for today
	emergencyBoard, err := s.GetEmergencyBoard()
	if err != nil {
		return nil, err
	}

	log.Printf("Copying emergency board to today: %s", emergencyBoard)
	bytes, err := os.ReadFile(s.GetBoardPath(emergencyBoard))
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(s.GetBoardPath(today), bytes, 0644); err != nil {
		return nil, err
	}

	res, err = s.GetTodaysBoard(ctx, today)
	if err == nil {
		return res, nil
	}

	err = fmt.Errorf("failed to serve todays board after writing emergency file: %w", err)
	panic(err)
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
	score := board.ScoreBoard(words)
	newBoard, err := board.PrintBoardJson(words)
	if err != nil {
		return nil, err
	}
	return &pb.WordsForBoardResponse{
		Board: newBoard,
		Score: float32(score),
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
