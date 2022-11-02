package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/marcoslhc/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse {
				Message: res,
			})

		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		log.Printf("received %s", req)
		res += fmt.Sprintf("Hello %s!\n", req.Name)
	}
}