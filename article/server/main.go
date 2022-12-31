package main

import (
	"log"
	"net"
    "context"
	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	pb "uditsharma.com/grpc-go-module/article/proto"
)

type Server struct {
	pb.ArticleServiceServer
}

var collection *mongo.Collection

var address string = "0.0.0.0:50051"

func main() {
    
	client , err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	
	if err!=nil{
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err !=nil{
		log.Fatal(err)
	}
	
	collection = client.Database("articledb").Collection("article")
	
	lis, err := net.Listen("tcp", address)

	log.Printf("Listening on %s\n", address)

	if err != nil {
		log.Fatalln("unable to start the TCP server")
	}

	s := grpc.NewServer()
   	pb.RegisterArticleServiceServer(s,&Server{});
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v\n", err)
	}
}
