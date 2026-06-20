package main

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "maps.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations complete")
}
