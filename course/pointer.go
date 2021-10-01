package main

import "log"

func main() {
	word := "Green"

	log.Println("The word is", word)
	changeUsingPointer(&word) // 把word的記憶體位址傳給function
	log.Println("After func call the word is", word)
}

// function接收某記憶體位址的值，而不會接一般的值
func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
