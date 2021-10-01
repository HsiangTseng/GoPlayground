package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "Sean",
		LastName:  "Tseng",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	var myNewVar float32

	myNewVar = 3.14
}
