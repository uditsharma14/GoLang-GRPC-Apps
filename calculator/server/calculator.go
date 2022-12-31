package main

import (
	"io"
	"log"

	pb "uditsharma.com/grpc-go-module/calculator/proto"
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

func (s *Server3) GetAverage(stream pb.PrimeFactorService_GetAverageServer) error {
	log.Printf("GetAverage function was invoked with")
	var count uint64 = 0
	var sum uint64 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("unable to do the average")
		}

		log.Printf("Receiving req: %v\n", req)
		sum = sum + req.Numer
		count++
	}

	return stream.SendAndClose(&pb.CalculateAverageResponse{
		AverageValue: float32(sum) / float32(count),
	})
}
