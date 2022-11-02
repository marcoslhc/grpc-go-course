package main

import (
	"context"
	"log"

	pb "github.com/marcoslhc/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet Called with %v", req)
	name := req.GetName()
	res := pb.GreetResponse{}
	res.Message = "Hello " + name + "!"
	return &res, nil
}
