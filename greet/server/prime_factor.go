package main

import (
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func (s *Server3) GetPrimeFactors(in *pb.PrimeFactorRequest, stream pb.PrimeFactorService_GetPrimeFactorsServer) error {
	log.Printf("GetPrimeFactors function was invoked with %v", in.Number)
	var num = in.Number
	var k = uint64(2)

	for num > 1 {
		if num%k == 0 {
			stream.Send(&pb.PrimeFactorResponse{
				PrimeFactor: k,
			})
			num = num / k
		} else {
			k = k + 1
		}
	}
	return nil
}
