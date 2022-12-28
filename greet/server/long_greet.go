package main

import (
	"io"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func (*Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with")
	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("unable to do the long greet")
		}

		log.Printf("Receiving req: %v\n", req)
		res += "Hello " + req.FirstName + "!\n"
	}
}
