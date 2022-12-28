package main

import (
	"context"
	"io"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func getPrimeFactor(c pb.PrimeFactorServiceClient) {
	log.Println("getPrimeFactor was invoked")

	req := &pb.PrimeFactorRequest{
		Number: 120,
	}

	stream, err := c.GetPrimeFactors(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not get the prime factor: %v\n", err)
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
			log.Fatalf("unable to get the prime factor:%v", err)
		}

		log.Printf("Prime factor is %v", msg.PrimeFactor)

	}

}
