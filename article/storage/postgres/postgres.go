package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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

// NewDBStringFromConfig build database connection string from config file.
func NewDBStringFromConfig(config *viper.Viper) (string, error) {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.dbname"),
		config.GetString("database.sslMode"),
	), nil
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
	return &stdb, nil
}

func NewTestStorage(dbstring string, migrationDir string) (*StoreDB, func()) {
	db, teardown := MustNewDevelopmentDB(dbstring, migrationDir)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return &StoreDB{Db: db}, teardown
}
