package main

import (
	"log"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Server not available: %v", err)
	}

	client := pb.NewCalculatorClient(conn)

	doPrimes(client, 1000000)
}
