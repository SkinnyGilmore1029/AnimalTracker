package models

import (
	"database/sql"
	"fmt"
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

func GetKids(db *sql.DB) []Kid {
	rows, err := db.Query("SELECT id, age, phone_number, name FROM kids")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var kids []Kid
	for rows.Next() {
		var k Kid
		err := rows.Scan(&k.ID, &k.Age, &k.PhoneNumber, &k.Name)
		if err != nil {
			log.Fatal(err)
		}
		kids = append(kids, k)
	}
	return kids
}

func RemoveKid(db *sql.DB, kidID int) {
	// Remove the kid
	sqlStmt := `DELETE FROM kids WHERE id = ?`
	_, err := db.Exec(sqlStmt, kidID)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// Set the owner_id of animals to NULL instead of deleting them
	sqlStmt = `UPDATE animals SET owner_id = NULL WHERE owner_id = ?`
	_, err = db.Exec(sqlStmt, kidID)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func ShowAllKids(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age, phone_number FROM kids")
	if err != nil {
		log.Fatalf("Failed to query kids: %v", err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id, age, phone int
		var name string
		err := rows.Scan(&id, &name, &age, &phone)
		if err != nil {
			log.Fatalf("Failed to scan kid: %v", err)
		}
		fmt.Println("=====================KIDS============================")
		fmt.Printf("Kid ID: %d\nName: %s\nAge: %d\nPhone: %d\n", id, name, age, phone)
		fmt.Println("=====================================================")
		count++
	}

	if count == 0 {
		fmt.Println("No kids in the database.")
	} else {
		fmt.Printf("Total kids: %d\n", count)
	}
}
