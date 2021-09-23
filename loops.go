package main

import "log"

func main() {
	animals := []string{"dog", "fish", "cat", "horse", "cow"}

	// 應該類似php的foreach($animas as $animal)
	for _, animal := range animals {
		log.Println(animal)
	}

	// ------------------------------------------------------------------

	values := make(map[string]string)
	values["dog"] = "Fido"
	values["cat"] = "John"

	for i, value := range values {
		log.Println(i, value)
	}

}
