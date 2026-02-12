package gui

import (
	models "AnimalTracker/Models"
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func OpenWindow(db *sql.DB) {
	a := app.New()
	w := a.NewWindow("Animal Tracker")

	// Make a vertical container to hold all the labels
	content := container.NewVBox()

	// Get kids from DB and add them as labels
	kids := models.GetKids(db)
	if len(kids) == 0 {
		content.Add(widget.NewLabel("No kids in the database."))
	} else {
		for _, k := range kids {
			label := widget.NewLabel(fmt.Sprintf("Name: %s, Age: %d, Phone: %d", k.Name, k.Age, k.PhoneNumber))
			content.Add(label)
		}
	}

	// Set the window content to the container with all labels
	w.SetContent(content)
	w.ShowAndRun()
}
