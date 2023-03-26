package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/sgielen/vierkantle/pkg/database"
	"github.com/sgielen/vierkantle/pkg/database/gendb"
	pb "github.com/sgielen/vierkantle/pkg/proto"
)

func (s *vierkantleService) SubmitScore(ctx context.Context, req *pb.SubmitScoreRequest) (*pb.SubmitScoreResponse, error) {
	if req.BoardName == "" {
		// BoardName was added 2023-03-26. If the client doesn't send it, assume this value.
		// TODO: make this an error instead.
		req.BoardName = "2023-03-26"
	}
	var dbUserId sql.NullInt64
	ctx, err := grpcWebAuth(ctx)
	if err == nil {
		userId, _, err := AuthUser(ctx)
		if err == nil {
			dbUserId = sql.NullInt64{Valid: true, Int64: userId}
		}
	}

	var res pb.SubmitScoreResponse
	if err := database.RunRWTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		return q.SetScore(ctx, gendb.SetScoreParams{
			BoardName:   req.BoardName,
			AnonymousID: int64(req.AnonymousId),
			UserID:      dbUserId,
			TeamSize:    req.TeamSize,
			Words:       req.Words,
			Seconds:     req.Seconds,
		})
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *vierkantleService) GetScores(ctx context.Context, req *pb.GetScoresRequest) (*pb.GetScoresResponse, error) {
	if req.Amount > 1000 {
		return nil, fmt.Errorf("max amount is 1000")
	}

	if req.BoardName == "" {
		// BoardName was added 2023-03-26. If the client doesn't send it, assume this value.
		// TODO: make this an error instead.
		req.BoardName = "2023-03-26"
	}

	if int32(req.Index) < 0 {
		req.Index = 0
	}

	var res pb.GetScoresResponse
	if err := database.RunROTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		scores, err := q.GetScores(ctx, gendb.GetScoresParams{
			BoardName: req.BoardName,
			Limit:     int32(req.Amount),
			Offset:    int32(req.Index),
		})
		if err != nil {
			return err
		}
		if len(scores) == 0 {
			res.TotalScores = 0
			return nil
		}

		res.Scores = make(map[int32]*pb.GetScoresResponse_Score, len(scores))
		res.TotalScores = int32(scores[0].FullCount)
		for i, score := range scores {
			offset := int32(req.Index) + int32(i)
			res.Scores[offset] = &pb.GetScoresResponse_Score{
				Name:     score.Username.String,
				TeamSize: score.TeamSize,
				Words:    score.Words,
				Seconds:  score.Seconds,
			}
			if score.AnonymousID == int64(req.MyAnonymousId) {
				res.YourScore = &offset
			}
		}
		if req.MyAnonymousId != 0 && res.YourScore == nil {
			myScore, err := q.GetMyScore(ctx, gendb.GetMyScoreParams{
				BoardName:   req.BoardName,
				AnonymousID: int64(req.MyAnonymousId),
			})
			if err == nil {
				rank := int32(myScore.Rank)
				res.YourScore = &rank
			} else if err != pgx.ErrNoRows {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &res, nil
}
