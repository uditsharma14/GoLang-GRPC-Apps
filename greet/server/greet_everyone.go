package main

import (
	"io"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client strean :%v\n", err)
		}

		res := "Hello" + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending the data to client")
		}

	}
}
