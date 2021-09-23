package main

import "log"

func main() {
	word := "Green"

	log.Println("The word is", word)
	changeUsingPointer(&word)
	log.Println("After func call the word is", word)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
