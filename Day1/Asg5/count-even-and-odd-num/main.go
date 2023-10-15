package main

import "fmt"

func main() {
	numbers := []int{12, 48, 53, 16, 75, 61, 37, 4}
	fmt.Println("The slice given here is:", numbers)
	even, odd := 0, 0

	for _, i := range numbers {
		if i&1 == 0 {
			even += 1
		} else {
			odd += 1
		}

	}

	fmt.Println("Count of even numbers in slice is:", even)
	fmt.Println("Count of odd numbers in slice is:", odd)
}
