package main

import (
	"context"
	"log"

	pb "uditsharma.com/grpc-go-module/calculator/proto"
)

func doSum(c pb.GetSumServiceClient) {
	log.Println("Do sum was invoked")
	res, err := c.Dosum(context.Background(), &pb.GetSumRequest{
		Num1: 12,
		Num2: 13,
	})

	if err != nil {
		log.Fatalf("Could not calculte sum: %v\n", err)
	}

	log.Printf("Sum : %v \n", res.Result)

}
