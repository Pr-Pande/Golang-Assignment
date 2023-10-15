package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}


type Person struct {
	Name    string
	Address 
}

func main() {
	person := Person{
		Name: "Ajay",
		Address: Address{
			Street:  "100 feet road",
			City:    "Bangalore",
			ZipCode: "560016",
		},
	}

	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Address: %s, %s, %s\n", person.Street, person.City, person.ZipCode)
}