package main

import (
	"context"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func (s *Server2) Dosum(ctx context.Context, in *pb.GetSumRequest) (*pb.GetSumResponse, error) {
	log.Printf("Dosum function was invoked with %v%v\n", in.Num1, in.Num2)
	return &pb.GetSumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
