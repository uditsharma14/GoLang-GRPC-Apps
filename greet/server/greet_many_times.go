package main

import (
	"fmt"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with")

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s,number %d", in.FirstName, i+1)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
