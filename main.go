package main

import (
	helpers "AnimalTracker/Helpers"
	models "AnimalTracker/Models"
	"fmt"
)

func main() {
	fmt.Println("Starting Animal Tracker Application...")
	db := helpers.CreateDatabase()
	defer db.Close()
	models.CreateKidTable(db)

	// Example of adding a kid to the database
	newKid := models.Kid{
		Name:        "John Doe",
		Age:         10,
		PhoneNumber: 1234567890,
	}
	models.AddKid(db, newKid)

	fmt.Println("Kid added to the database successfully!")
}
