package main

import "fmt"

func main() {
	numbers := []int{12, 48, 53, 16, 75, 61, 37, 4}
	fmt.Println("The slice given here is:", numbers)

	for i := range numbers {
		numbers[i] *= 2
	}

	fmt.Println("The modified slice is:", numbers)
}
