package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "uditsharma.com/grpc-go-module/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

var address string = "0.0.0.0:50051"

func main() {

	lis, err := net.Listen("tcp", address)

	log.Printf("Listening on %s\n", address)

	if err != nil {
		log.Fatalln("unable to start the TCP server")
	}

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})


	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v\n", err)
	}
}
