package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Printf("Error occured while calling the greetManytimes %v", err)

	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Udit"},
		{FirstName: "Sumit"},
		{FirstName: "Mohit"},
	}

	waitc := make(chan chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request : %v\n", req)
			stream.Send(req)
			time.Sleep((1 * time.Second))
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while recieving : %v", err)
				break
			}

			log.Printf("Recieved : %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
