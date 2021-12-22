package main

import (
	"context"
	"fmt"
	"log"
	"practice/blog/gunk/v1/article"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing grpc: %s", err)
	}
	c := article.NewArticleClient(conn)
	res, err := c.GetArticle(context.Background(), &article.GetArticleRequest{
		ID: 1,
	})
	if err != nil {
		log.Fatal("error calling service ", err)
	}
	fmt.Printf("priniting response: %#v", res)
}
