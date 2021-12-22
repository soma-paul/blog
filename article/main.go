package main

import (
	"log"
	"net"

	tc "practice/blog/article/core/article"
	"practice/blog/article/services/article"
	"practice/blog/article/storage/postgres"
	apb "practice/blog/gunk/v1/article"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatal("failed to listen to the port 4444: ", err)
	}
	grpcServer := grpc.NewServer()
	st, _ := postgres.NewStorage(postgres.DbConfig())
	cs := tc.NewCoreSvc(st)
	s := article.NewArticleServer(cs)
	apb.RegisterArticleServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve ", err)
	}
}
