package main

import (
	"context"
	"log"

	pb "github.com/lordsonvimal/golang-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstOperand:  2,
		SecondOperand: 3,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
