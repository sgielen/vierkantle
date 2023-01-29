package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	pb "github.com/sgielen/vierkantle/pkg/proto"
)

// TODO: this should be moved into a database instead of being in-memory.
var mtx sync.Mutex
var scores []*pb.SubmitScoreRequest

// TODO: boards should have a name. Pass this name in SubmitScore and GetScores
// so we don't need to wipe anything.
var lastActionTime time.Time

func resetScoresIfNecessary() {
	now := time.Now()

	if location, err := time.LoadLocation("Europe/Amsterdam"); err == nil {
		now = now.In(location)
	}

	if lastActionTime.IsZero() {
		lastActionTime = now
		return
	}

	// If the last action time was before 12:00 pm Europe/Amsterdam time,
	// and the new action time is after, return true.
	if lastActionTime.Hour() < 12 && now.Hour() > 12 {
		scores = nil
	}
	lastActionTime = now
}

func (s *vierkantleService) SubmitScore(ctx context.Context, req *pb.SubmitScoreRequest) (*pb.SubmitScoreResponse, error) {
	mtx.Lock()
	defer mtx.Unlock()

	resetScoresIfNecessary()

	found := false
	for i, s := range scores {
		if s.AnonymousId == req.AnonymousId {
			scores[i] = req
			found = true
			break
		}
	}
	if !found {
		scores = append(scores, req)
	}

	sort.SliceStable(scores, func(i, j int) bool {
		onWords := scores[j].Words - scores[i].Words
		if onWords != 0 {
			return onWords < 0
		}
		return scores[j].Seconds > scores[i].Seconds
	})

	return &pb.SubmitScoreResponse{}, nil
}

func (s *vierkantleService) GetScores(ctx context.Context, req *pb.GetScoresRequest) (*pb.GetScoresResponse, error) {
	mtx.Lock()
	defer mtx.Unlock()

	resetScoresIfNecessary()

	if req.Amount > 1000 {
		return nil, fmt.Errorf("max amount is 1000")
	}

	// Note: req.Index may be >= len(scores)
	end := int32(req.Index + req.Amount)
	if end > int32(len(scores)) {
		end = int32(len(scores))
	}

	res := &pb.GetScoresResponse{}
	res.Scores = make(map[int32]*pb.GetScoresResponse_Score, req.Amount)
	res.TotalScores = int32(len(scores))
	for i := int32(req.Index); i < end; i += 1 {
		score := scores[i]
		res.Scores[i] = &pb.GetScoresResponse_Score{
			Name:     "",
			TeamSize: score.TeamSize,
			Words:    score.Words,
			Seconds:  score.Seconds,
		}
		if req.MyAnonymousId != 0 && score.AnonymousId == int32(req.MyAnonymousId) {
			res.YourScore = i
		}
	}

	if req.MyAnonymousId != 0 && res.YourScore == 0 {
		for i, s := range scores {
			if s.AnonymousId == int32(req.MyAnonymousId) {
				res.YourScore = int32(i)
				break
			}
		}
	}

	return res, nil
}
