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

type Server2 struct {
	pb.GetSumServiceServer
}

type Server3 struct {
	pb.PrimeFactorServiceServer
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
	pb.RegisterGetSumServiceServer(s, &Server2{})
	pb.RegisterPrimeFactorServiceServer(s, &Server3{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v\n", err)
	}
}
