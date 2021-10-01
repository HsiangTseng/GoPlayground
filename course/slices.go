package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	// 類似array_push
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 6)
	mySlice = append(mySlice, 5)

	sort.Ints(mySlice)

	log.Println(mySlice)

	// ----------------------------------------------------

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers[5:9])

	// ----------------------------------------------------

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
