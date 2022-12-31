package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "uditsharma.com/grpc-go-module/greet/proto"
)

var address string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	defer conn.Close()

	doGreetManyTimes(c)
	doLongGreet(c)
	
	doGreetEveryone(c)

}
