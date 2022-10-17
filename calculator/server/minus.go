package main

import (
	"io"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/calculator/proto"
)

func (s *Server) Minus(stream pb.CalculatorService_MinusServer) error {
	log.Println("Minus function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		err = stream.Send(&pb.BinaryResponse{
			Result: req.FirstOperand - req.SecondOperand,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
