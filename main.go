package main

import (
	gui "AnimalTracker/GUI"
	helpers "AnimalTracker/Helpers"
	"fmt"
)

func main() {
	fmt.Println("Starting Animal Tracker Application...")
	db := helpers.CreateDatabase()
	defer db.Close()
	gui.OpenWindow()
}
