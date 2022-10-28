package main

import (
	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.Calculator_PrimesServer) error {
	var num = req.GetNumber()
	var k int64 = 2
	var res []int64
	for {
		if num <= 1 {
			break
		}

		if num%k == 0 {
			res = append(res, k)
			stream.Send(&pb.PrimeResponse{
				Result: res,
			})
			num = num / k
		} else {
			k = k + 1
		}
	}
	return nil
}
