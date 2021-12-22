package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type StoreDB struct {
	Db *sqlx.DB
}

const DRIVER = "postgres"

func DbConfig() string {
	dbParams := " " + "user=postgres"
	dbParams += " " + "host=localhost"
	dbParams += " " + "dbname=blog"
	dbParams += " " + "password=password"
	dbParams += " " + "sslmode=disable"

	return dbParams
}

func NewStorage(dbconfig string) (*StoreDB, error) {
	db, err := sqlx.Connect(DRIVER, dbconfig)
	if err != nil {
		log.Fatal("error connecting to db:", err)
	}
	stdb := StoreDB{
		Db: db,
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return &stdb, err
}
