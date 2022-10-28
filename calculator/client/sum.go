package main

import (
	"context"
	"log"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorClient, a, b int64) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		First: a,
		Second: b,
	})

	if err != nil {
		log.Fatalf("Error doing request: %v", err)
	}

	log.Printf("%v", res)
}