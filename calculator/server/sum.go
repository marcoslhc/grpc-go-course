package main

import (
	"context"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (res *pb.SumResponse, err error) {
	res = &pb.SumResponse{
		Sum: req.GetFirst() + req.GetSecond(),
	}
	return res, nil
}