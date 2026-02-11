package models

type Animal struct {
	ID      int
	Name    string
	Gender  string
	Breed   string
	Age     int
	OwnerID int
	Picture []byte
}
