package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lordsonvimal/golang-grpc/greet/proto"
)

func doGreetClientStream(c pb.GreetServiceClient) {
	log.Println("doGreetClientStream was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Lordson"},
		{FirstName: "Vimal"},
		{FirstName: "Logesh"},
		{FirstName: "Deepan"},
	}

	stream, err := c.GreetClientStream(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetClientStream: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from GreetClientStream: %v\n", err)
	}

	log.Printf("GreetClientStream: %s\n", res.Result)
}
