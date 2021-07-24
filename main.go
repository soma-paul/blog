package main

import (
	"net/http"

	"Blog/handler"
	"Blog/storage/postgres"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

func dbConfig() string {
	dbParams := " " + "user=postgres"
	dbParams += " " + "host=localhost"
	dbParams += " " + "dbname=blog"
	dbParams += " " + "password=test123"
	dbParams += " " + "sslmode=disable"

	return dbParams
}

func main() {
	store, err := postgres.NewStorage(dbConfig()) //establish database connection
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
