package main

type Person struct {
	FirstName string `json:"first_name`
	LastName  string `json:"last_name`
	HairColor string `json:"hair`
	HasDog    bool   `json:"has_dog`
}

func main() {
	myJson := `
	[
		{
			"first_name" : "Clark",
			"last_name" : "Kent,
			"hair" : "black",
			"has_dog" : false
		},
		{
			"first_name" : "Bruce",
			"last_name" : "Wayne,
			"hair" : "black",
			"has_dog" : false
		}
	]
	`

	//marshalled = 整理、整齊
	var unmarshalled []Person

	err := json.unmarshalled([]byte(myJson), &unmarshalled)
}
