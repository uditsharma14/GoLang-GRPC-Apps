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

var address string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln("unable to start the TCP server")
	}

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterGetSumServiceServer(s, &Server2{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v\n", err)
	}

	log.Fatalf("Listening on %s\n", address)

}
