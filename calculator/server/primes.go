package main

import (
	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.Calculator_PrimesServer) error {
	var num = req.GetNumber()
	var k int64 = 2

	for num > 1 {
		if num%k == 0 {
			err := stream.Send(&pb.PrimeResponse{
				Result: k,
			})

			if err != nil {
				return err
			}
			num = num / k
		} else {
			k = k + 1
		}
	}
	return nil
}
