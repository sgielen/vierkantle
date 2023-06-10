package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/sgielen/vierkantle/pkg/database"
	"github.com/sgielen/vierkantle/pkg/database/gendb"
	"github.com/sgielen/vierkantle/pkg/email"
	pb "github.com/sgielen/vierkantle/pkg/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *vierkantleService) AddBoardToQueue(ctx context.Context, req *pb.AddBoardToQueueRequest) (*pb.AddBoardToQueueResponse, error) {
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, username, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}

	var res pb.AddBoardToQueueResponse
	if err := database.RunRWTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		id, err := q.AddBoardToQueue(ctx, gendb.AddBoardToQueueParams{
			UserID:    userId,
			Board:     req.Board,
			BoardName: req.BoardName,
		})
		if err != nil {
			return err
		}
		res.Id = id

		return email.SendNewBoardInQueue(ctx, username, req.BoardName)
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *vierkantleService) ListBoardQueue(ctx context.Context, req *pb.ListBoardQueueRequest) (*pb.ListBoardQueueResponse, error) {
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, _, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}

	var res pb.ListBoardQueueResponse
	if err := database.RunROTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		user, err := q.GetUserById(ctx, userId)
		if err != nil {
			return err
		}

		if !user.IsAdmin {
			return fmt.Errorf("must be admin")
		}

		boards, err := q.ListBoardQueue(ctx)
		if err != nil {
			return err
		}

		for _, board := range boards {
			res.Boards = append(res.Boards, &pb.ListBoardQueueResponse_Board{
				Id:        board.ID,
				UserId:    board.UserID,
				Username:  board.Username,
				BoardName: board.BoardName,
				AddedAt:   timestamppb.New(board.AddedAt),
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *vierkantleService) GetBoardFromQueue(ctx context.Context, req *pb.GetBoardFromQueueRequest) (*pb.GetBoardFromQueueResponse, error) {
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, _, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}

	var res pb.GetBoardFromQueueResponse
	if err := database.RunROTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		user, err := q.GetUserById(ctx, userId)
		if err != nil {
			return err
		}

		if !user.IsAdmin {
			return fmt.Errorf("must be admin")
		}

		board, err := q.GetBoardFromQueue(ctx, req.Id)
		if err != nil {
			return err
		}

		res.Board = board.Board
		res.BoardName = board.BoardName

		return nil
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *vierkantleService) UpdateBoardInQueue(ctx context.Context, req *pb.UpdateBoardInQueueRequest) (*pb.UpdateBoardInQueueResponse, error) {
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, _, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}

	var res pb.UpdateBoardInQueueResponse
	if err := database.RunRWTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		user, err := q.GetUserById(ctx, userId)
		if err != nil {
			return err
		}

		if !user.IsAdmin {
			return fmt.Errorf("must be admin")
		}

		return q.UpdateBoardInQueue(ctx, gendb.UpdateBoardInQueueParams{
			ID:        req.Id,
			Board:     req.Board,
			BoardName: req.BoardName,
		})
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *vierkantleService) RemoveBoardsFromQueue(ctx context.Context, req *pb.RemoveBoardsFromQueueRequest) (*pb.RemoveBoardsFromQueueResponse, error) {
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, _, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}

	var res pb.RemoveBoardsFromQueueResponse
	if err := database.RunRWTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		user, err := q.GetUserById(ctx, userId)
		if err != nil {
			return err
		}

		if !user.IsAdmin {
			return fmt.Errorf("must be admin")
		}

		return q.RemoveBoardsFromQueue(ctx, req.Ids)
	}); err != nil {
		return nil, err
	}

	return &res, nil
}
