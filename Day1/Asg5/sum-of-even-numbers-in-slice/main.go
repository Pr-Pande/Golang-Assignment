package main

import "fmt"

func main() {
	numbers := []int{5, 6, 3, 9, 12, 7, 24, 14}
	fmt.Println("The slice given here is:", numbers)
	even := 0

	for _, i := range numbers {
		if i&1 == 0 {
			even += i
		}
	}

	fmt.Println("Sum of all even numbers in slice is:", even)

}
