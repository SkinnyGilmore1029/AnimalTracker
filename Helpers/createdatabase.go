package helpers

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./AnimalTracker.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DataBase Is Ready")
	return db
}
