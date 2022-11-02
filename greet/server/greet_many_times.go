package main

import (
	"fmt"

	pb "github.com/marcoslhc/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s!, number: %d", req.Name, i)
		message := &pb.GreetResponse{
			Message: res,
		}
		err := stream.Send(message)
		if err != nil {
			return err
		}
	}
	return nil
}
