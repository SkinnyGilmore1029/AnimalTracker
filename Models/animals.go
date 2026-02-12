package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Animal struct {
	ID      int
	Name    string
	Gender  string
	Breed   string
	Age     int
	OwnerID int
	Picture []byte
}

func CreateAnimalTable(db *sql.DB) {
	sqlStmt := `CREATE TABLE IF NOT EXISTS animals (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	gender TEXT,
	breed TEXT,
	age INTEGER,
	owner_id INTEGER,
	picture BLOB,
	FOREIGN KEY(owner_id) REFERENCES kids(id)
);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func AddAnimal(db *sql.DB, animal Animal) {
	sqlStmt := `INSERT INTO animals (name, gender, breed, age, owner_id, picture) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, animal.Name, animal.Gender, animal.Breed, animal.Age, animal.OwnerID, animal.Picture)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func GetAnimals(db *sql.DB) []Animal {
	rows, err := db.Query("SELECT id, name, gender, breed, age, owner_id, picture FROM animals")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var a Animal
		err := rows.Scan(&a.ID, &a.Name, &a.Gender, &a.Breed, &a.Age, &a.OwnerID, &a.Picture)
		if err != nil {
			log.Fatal(err)
		}
		animals = append(animals, a)
	}
	return animals
}

func RemoveAnimal(db *sql.DB, animalID int) {
	sqlStmt := `DELETE FROM animals WHERE id = ?`
	_, err := db.Exec(sqlStmt, animalID)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func ShowAllAnimals(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, gender, breed, age, owner_id FROM animals")
	if err != nil {
		log.Fatalf("Failed to query animals: %v", err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id, age, ownerID sql.NullInt64
		var name, gender, breed string
		err := rows.Scan(&id, &name, &gender, &breed, &age, &ownerID)
		if err != nil {
			log.Fatalf("Failed to scan animal: %v", err)
		}

		owner := "None"
		if ownerID.Valid {
			owner = fmt.Sprintf("%d", ownerID.Int64)
		}
		fmt.Println("=============ANIMALS===================================")
		fmt.Printf("Animal ID: %d\nName: %s\nGender: %s\nBreed: %s\nAge: %d\nOwnerID: %s\n",
			id.Int64, name, gender, breed, age.Int64, owner)
		fmt.Println("=====================================================")
		count++
	}

	if count == 0 {
		fmt.Println("No animals in the database.")
	} else {
		fmt.Printf("Total animals: %d\n", count)
	}
}
