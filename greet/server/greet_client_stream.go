package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/greet/proto"
)

func (s *Server) GreetClientStream(stream pb.GreetService_GreetClientStreamServer) error {
	log.Println("GreetClientStream function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
