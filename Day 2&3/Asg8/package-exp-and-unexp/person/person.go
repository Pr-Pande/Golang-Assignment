package person

import (
	"fmt"
	"package-exp-and-unexp/models"
)

func PrintPerson(p models.Person) {
	fmt.Println("Name of person:", p.Name)
	fmt.Println("\nAge:", p.Age)
}