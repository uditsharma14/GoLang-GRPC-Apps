package main

import (
	"context"
	"log"

	pb "uditsharma.com/grpc-go-module/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Udit",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting : %s \n", res.Result)

}
