package main

import(
	pb "uditsharma.com/grpc-go-module/article/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ArticleItem struct{
	ID primitive.ObjectID ``
	AuthorId string ``
	Title string ``
	Content string ``
}

func documentToBlog(data *ArticleItem) *pb.Article{
     return &pb.Article{
		 Id : data.ID.Hex(),
		 AuthorId : data.AuthorId,
		 Title : data.Title,
		 Content : data.Content,
	 }
}