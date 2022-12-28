package main

import (
	"context"
	"io"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do greetmanytimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Udit",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not greet many times: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			if msg == nil {
				log.Fatalf("failed")
			}
			log.Fatalf("unable to get the greet many times %v", err)
		}

		log.Printf("message recieved from greet many times %s", msg.Result)

	}

}
