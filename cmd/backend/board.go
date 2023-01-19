package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	pb "github.com/sgielen/vierkantle/pkg/proto"
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
