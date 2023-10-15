package main

import "fmt"

// Define a student struct to store student information
type student struct {
	Age   int
	Grade float64
	City  string
}

func main() {
	studentData := make(map[string]student)

	// Add student information to the studentData map
	studentData["Varun"] = student{Age: 18, Grade: 9.1, City: "Delhi"}
	studentData["Sagar"] = student{Age: 17, Grade: 8.5, City: "Jaipur"}
	studentData["Nikhil"] = student{Age: 19, Grade: 7.6, City: "Mumbai"}

	// Retrieve and print the details of each student
	for name, data := range studentData {
		fmt.Printf("Student: %s\n", name)
		fmt.Printf("Age: %v\n", data.Age)
		fmt.Printf("Grade: %g\n", data.Grade)
		fmt.Printf("City: %s\n", data.City)
		fmt.Println()
	}
}
