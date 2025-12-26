package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	database *sql.DB
)

func init() {
	var err error
	database, err = sql.Open("sqlite3", "./pricing.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to sqlite")
}
