package main

import "fmt"

func main() {
	var numbers []int

	numbers = append(numbers, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)
	numbers = append(numbers, 25)
	fmt.Println(numbers)

	fmt.Printf("len=%d cap=%d\n", len(numbers), cap(numbers))
	numbers = remove(numbers, 2)
	fmt.Println(numbers)
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}