package main

import (
	"context"
	"log"
	"time"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func (e *executor) doAverage() {
	nums := []*pb.AverageRequest{
		{Number: 1.0},
		{Number: 2.0},
		{Number: 3.0},
		{Number: 4.0},
	}

	stream, err := e.client.Average(context.Background())

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	for _, num := range nums {
		stream.Send(num)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Printf("%v\n", res)
}
