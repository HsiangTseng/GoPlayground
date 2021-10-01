package main

import "fmt"

type Animal interface {
	Says() string
	NumberOFLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 39,
	}

	PrintInfo(&dog)
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "and has", a.NumberOFLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOFLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOFLegs() int {
	return 2
}
