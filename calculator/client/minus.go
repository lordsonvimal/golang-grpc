package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/lordsonvimal/golang-grpc/calculator/proto"
)

func doMinus(c pb.CalculatorServiceClient) {
	log.Println("doMinus was invoked")

	stream, err := c.Minus(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.BinaryRequest{
		{FirstOperand: 2, SecondOperand: 1},
		{FirstOperand: 10, SecondOperand: 20},
		{FirstOperand: 13, SecondOperand: 4},
		{FirstOperand: 7, SecondOperand: 5},
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

			log.Printf("doMinus from server: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
