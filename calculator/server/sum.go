package main

import (
	"context"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.BinaryRequest) (*pb.BinaryResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.BinaryResponse{
		Result: in.FirstOperand + in.SecondOperand,
	}, nil
}
