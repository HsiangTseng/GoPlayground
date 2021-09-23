package main

import (
	"log"
	"time"
)

var s = "seven"

//
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Sean",
		LastName:    "Tseng",
		PhoneNumber: "0910313342",
	}

	log.Println(user.FirstName, user.LastName, "PhoneNumber:", user.PhoneNumber)
}
