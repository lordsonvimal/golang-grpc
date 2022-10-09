package main

import (
	"io"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/greet/proto"
)

func (s *Server) GreetBiDirectionStream(stream pb.GreetService_GreetBiDirectionStreamServer) error {
	log.Println("GreetBiDirectionStream function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
