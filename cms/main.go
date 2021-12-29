package main

import (
	"fmt"
	"log"
	"net/http"
	"practice/blog/article/storage/postgres"
	apb "practice/blog/gunk/v1/article"
	"practice/blog/handler"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
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

	store, err := newDBFromConfig(config) //establish database connection
	if err != nil {
		log.Println("database connection problem: ", err)
	}

	decoder := schema.NewDecoder()
	session := sessions.NewCookieStore([]byte("6ret-key"))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.GetString("article.host"), config.GetString("article.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing grpc: %s", err)
	}

	c := apb.NewArticleClient(conn)
	r, err := handler.NewServer(store, decoder, session, c)
	if err != nil {
		log.Println("error creating newServer: ", err)
	}
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", config.GetString("server.host"), config.GetString("server.port")),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("error listening and serving: ", err)
	}
	// res, err := c.GetArticle(context.Background(), &article.GetArticleRequest{
	// 	ID: 1,
	// })
	// if err != nil {
	// 	log.Fatal("error calling service ", err)
	// }
	// fmt.Printf("priniting response: %#v", res)
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
