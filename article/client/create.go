package main

import (
	"log"
	 pb "uditsharma.com/grpc-go-module/article/proto"
	 "context"
)


func CreateArticle(c pb.ArticleServiceClient) string {
	log.Println("--creating article was started")

   article := &pb.Article{
	   AuthorId : "Udit",
	   Title : "My First article",
	   Content : "Content of first article is always precious",
   }
   res , err := c.CreateArticle(context.Background(),article)
   
 
   if err !=nil{
      log.Fatal("Unable to call the create method")
   }
   log.Printf("Article has been created successfully %v",res.Id)
   return res.Id;
}