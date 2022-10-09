package main

import (
	"fmt"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/greet/proto"
)

func (s *Server) GreetStream(in *pb.GreetRequest, stream pb.GreetService_GreetStreamServer) error {
	log.Printf("GreetStream function was invoked with: %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
