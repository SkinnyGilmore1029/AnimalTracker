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
	models.CreateAnimalTable(db)

	models.ShowAllKids(db)
	models.ShowAllAnimals(db)

	fmt.Println("=====================================================")
	fmt.Println("Database setup complete. You can now add kids and animals to the tracker.")
}
