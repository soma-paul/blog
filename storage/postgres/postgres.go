package postgres

import "github.com/jmoiron/sqlx"

type StoreDB struct {
	Db *sqlx.DB
}

const DRIVER = "postgres"

func NewStorage(dbconfig string) (*StoreDB, error) {
	db, err := sqlx.Connect(DRIVER, dbconfig)
	stdb := StoreDB{
		Db: db,
	}
	return &stdb, err
}
