package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func incrementAge(s *Student) {
	s.Age++
}

func main() {
	student := Student{
		Name: "John",
		Age:  20,
	}

	fmt.Printf("Name: %s, Age: %d\n", student.Name, student.Age)
	incrementAge(&student)

	fmt.Printf("Updated Age for %s is %d\n", student.Name, student.Age)
}