package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
	Name string
}

func (d Dog) MakeSound() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) MakeSound() string {
	return "Meow!"
}

func main() {
	dog := Dog{Name: "Tommy"}
	cat := Cat{Name: "Lucy"}

	fmt.Printf("%s says: %s\n", dog.Name, dog.MakeSound())
	fmt.Printf("%s says: %s\n", cat.Name, cat.MakeSound())
}