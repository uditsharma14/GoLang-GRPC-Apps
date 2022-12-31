package main

import (
	pb "uditsharma.com/grpc-go-module/article/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s* Server) CreateArticle(ctx context.Context, in *pb.Article,) (*pb.ArticleId, error){
 log.Printf("Create Article was invoked with %v\n",in)

 data := ArticleItem{
	AuthorId : in.AuthorId,
	Title : in.Title,
	Content : in.Content,
 }

   res,err :=  collection.InsertOne(ctx,data)

 if err!= nil{
    return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal Error : %v\n",err),
	)
 }

   oid , ok := res.InsertedID.(primitive.ObjectID)


   if !ok{
	   return nil, status.Errorf(
		   codes.Internal,
		   "Cannot convert to OID",
	   )
   }

   return &pb.ArticleId{
	   Id : oid.Hex(),
   },nil

}
 