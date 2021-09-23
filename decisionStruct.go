package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// -------------------------------------------
	cat := "cat"
	if cat == "cat" {
		log.Println("CAT!")
	} else {
		log.Println("NO CAT!")
	}

	// --------------------------------------------
	num := 80
	is := true

	if num > 99 && is {
		log.Println("All True")
	} else {
		log.Println("Someone False")
	}

	// ---------------------------------------------
	myVar := "56w"
	switch myVar {
	case "cat":
		log.Println("Case cat")
	case "dog":
		log.Println("Case dog")
	case "fish":
		log.Println("Case fish")
	default:
		log.Println("Default")
	}

}
