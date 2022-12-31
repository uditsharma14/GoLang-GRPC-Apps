package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "uditsharma.com/grpc-go-module/calculator/proto"
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

func doCalculateAverage(c pb.PrimeFactorServiceClient) {
	log.Println("doCalculateAverage was invoked")

	reqs := []*pb.CalculteAverageRequest{
		{Numer: 23},
		{Numer: 25},
		{Numer: 27},
		{Numer: 29},
		{Numer: 35},
	}

	stream, err := c.GetAverage(context.Background())

	if err != nil {
		log.Fatalf("Could not calculate average: %v\n", err)
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

	log.Printf("Result recieved from server %v", res.AverageValue)

}
