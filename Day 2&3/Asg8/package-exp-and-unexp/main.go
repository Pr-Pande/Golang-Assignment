package main

import (
	"package-exp-and-unexp/models"
	"package-exp-and-unexp/person"
)

func main() {
	p := models.Person{
		Name: "Vinay",
		Age: 25,
	}

	person.PrintPerson(p)
}