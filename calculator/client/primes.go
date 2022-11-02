package main

import (
	"context"
	"io"
	"log"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func (e *executor) doPrimes(n int64) {
	res, err := e.client.Primes(context.Background(), &pb.PrimeRequest{
		Number: n,
	})

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading the stream %v", err)
		}

		log.Printf("%v\n", msg)
	}
}
