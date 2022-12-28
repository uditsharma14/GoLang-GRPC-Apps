package main

import (
	"context"
	"log"
	"time"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Udit"},
		{FirstName: "Mohit"},
		{FirstName: "Suresh"},
		{FirstName: "Sunil"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Could not greet long: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error While Recieving the data from stream: %v\n", err)
	}

	log.Printf("Result recieved from server %v", res.Result)
}
