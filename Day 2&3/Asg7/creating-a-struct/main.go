package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	emp1 := Employee{
		Id:   1,
		Name: "Kunal",
		Age:  28,
		City: "Delhi",
	}

	emp2 := Employee{
		Id:   2,
		Name: "Mehul",
		Age:  35,
		City: "Kochi",
	}

	fmt.Printf("%+v\n", emp1)
	fmt.Printf("%+v\n", emp2)
}
