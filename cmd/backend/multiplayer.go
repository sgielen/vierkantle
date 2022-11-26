package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/sgielen/vierkantle/pkg/proto"
)

type vierkantleService struct{}

func (s *vierkantleService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("hello called")
	if req.M == "world" {
		return &pb.HelloResponse{
			M: fmt.Sprintf("Hello %s indeed!", req.M),
		}, nil
	}
	return nil, fmt.Errorf("unimplemented")
}
func (s *vierkantleService) HelloStream(stream pb.VierkantleService_HelloStreamServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.Send(&pb.HelloStreamResponse{
				M: "Okay bye!",
			})
		}
		if err := stream.Send(&pb.HelloStreamResponse{
			M: fmt.Sprintf("A hello %s to you as well!", msg.M),
		}); err != nil {
			return err
		}
	}
}
