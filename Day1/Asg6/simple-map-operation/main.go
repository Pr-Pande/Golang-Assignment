package main

import "fmt"

func main() {
	studentGrades := make(map[string]float64)
	studentGrades["Ajay"] = 9.8
	studentGrades["Sonu"] = 8.4
	studentGrades["Sagar"] = 7.2

	fmt.Println(studentGrades)
	delete(studentGrades, "Sonu")
	fmt.Println("After deleting Sonu's entry from the map... \n")
	fmt.Println(studentGrades)

}
