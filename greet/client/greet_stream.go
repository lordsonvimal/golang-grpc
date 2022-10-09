package main

import (
	"context"
	"io"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/greet/proto"
)

func doGreetStream(c pb.GreetServiceClient) {
	log.Println("doGreetStream was invoked")
	req := &pb.GreetRequest{
		FirstName: "Lordson",
	}

	stream, err := c.GreetStream(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}

		log.Printf("Greeting: %s\n", msg.Result)
	}
}
