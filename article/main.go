package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	tc "practice/blog/article/core/article"
	"practice/blog/article/services/article"
	"practice/blog/article/storage/postgres"
	apb "practice/blog/gunk/v1/article"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("error loading configuration: %v", err)
	}

	grpcServer := grpc.NewServer()
	st, err := newDBFromConfig(config)
	if err != nil {
		log.Fatal("error creating db connection from config: ", err)
	}
	cs := tc.NewCoreSvc(st)
	s := article.NewArticleServer(cs)
	apb.RegisterArticleServer(grpcServer, s)

	host, port := config.GetString("server.host"), config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal("failed to listen to the port 4444: ", err)
	}
	// reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve ", err)
	}
}

func newDBFromConfig(config *viper.Viper) (*postgres.StoreDB, error) {
	cf := func(c string) string { return config.GetString("database." + c) }
	ci := func(c string) string { return strconv.Itoa(config.GetInt("database." + c)) }
	dbParams := " " + "user=" + cf("user")
	dbParams += " " + "host=" + cf("host")
	dbParams += " " + "port=" + cf("port")
	dbParams += " " + "dbname=" + cf("dbname")
	if password := cf("password"); password != "" {
		dbParams += " " + "password=" + password
	}
	dbParams += " " + "sslmode=" + cf("sslMode")
	dbParams += " " + "connect_timeout=" + ci("connectionTimeout")
	dbParams += " " + "statement_timeout=" + ci("statementTimeout")
	dbParams += " " + "idle_in_transaction_session_timeout=" + ci("idleTransacionTimeout")
	db, err := postgres.NewStorage(dbParams)
	if err != nil {
		return nil, err
	}
	return db, nil
}
