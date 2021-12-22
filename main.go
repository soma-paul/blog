package main

import (
	"net/http"

	"practice/blog/article/storage/postgres"
	"practice/blog/handler"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

func main() {
	store, err := postgres.NewStorage(postgres.DbConfig()) //establish database connection
	handler.CheckError("database connection problem: ", err)

	decoder := schema.NewDecoder()
	session := sessions.NewCookieStore([]byte("6ret-key"))

	r, err := handler.NewServer(store, decoder, session)
	handler.CheckError("error on creating new server: ", err)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	err = srv.ListenAndServe()
	handler.CheckError("error listening and serving: ", err)

}
