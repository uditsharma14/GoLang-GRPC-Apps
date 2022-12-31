package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "uditsharma.com/grpc-go-module/calculator/proto"
)

var address string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}



	d := pb.NewPrimeFactorServiceClient(conn)
    e := pb.NewGetSumServiceClient(conn)
	

	defer conn.Close()
	getPrimeFactor(d)
	doSum(e)
	doCalculateAverage(d)


}
