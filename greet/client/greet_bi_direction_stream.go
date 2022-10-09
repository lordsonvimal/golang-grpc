package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/lordsonvimal/golang-grpc/greet/proto"
)

func doBiDirectionStream(c pb.GreetServiceClient) {
	log.Println("doBiDirectionStream was invoked")

	stream, err := c.GreetBiDirectionStream(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Lordson"},
		{FirstName: "Vimal"},
		{FirstName: "Logesh"},
		{FirstName: "Deepan"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("GreetClientStream: %s\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
