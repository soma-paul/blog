package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations/sql", "directory with migration files")
)

func main() {
	flags.Usage = usage
	_ = flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	switch args[0] {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case "-h", "--help":
		flags.Usage()
		return
	}

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]
	driver := "postgres"
	dbString := newDBFromConfig()

	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	if dbString == "" {
		log.Fatalf("-dbString=%q not supported\n", dbString)
	}

	db, err := sql.Open(driver, dbString)
	if err != nil {
		log.Fatalf("Invalid DB string:%q %v\n", dbString, err)
	}

	defer db.Close()
	var arguments []string
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

func newDBFromConfig() string {
	dbParams := " " + "user=postgres"
	//dbParams += " " + "host=localhost"
	//dbParams += " " + "port=5435"
	dbParams += " " + "dbname=blog"
	dbParams += " " + "password=test123"
	dbParams += " " + "sslmode=disable"

	return dbParams
}

func usage() {
	log.Print(usageCommands)
}

var (
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
	fix                  Apply sequential ordering to migrations
`
)
