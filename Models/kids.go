package models

import (
	"database/sql"
	"log"
)

type Kid struct {
	ID          int
	Name        string
	Age         int
	PhoneNumber int
}

func CreateKidTable(db *sql.DB) {
	sqlStmt := `CREATE TABLE IF NOT EXISTS kids (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age INTEGER,
	phone_number INTEGER
);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func AddKid(db *sql.DB, kid Kid) {
	sqlStmt := `INSERT INTO kids (name, age, phone_number) VALUES (?, ?, ?)`
	_, err := db.Exec(sqlStmt, kid.Name, kid.Age, kid.PhoneNumber)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}
