package main

import (
	"io"
	"log"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func (s *Server) Average(stream pb.Calculator_AverageServer) error {
	res := float32(0);
	sum := float32(0);
	n := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: res,
			})

		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		n++
		sum += req.Number
		res = sum / float32(n)
		log.Printf("received %-4s, n: %-4d, sum: %-4f, res: %-4f", req, n, sum, res)

	}
}