package main

import "fmt"

func main() {
	numbers := []int{75, 42, 63, 8, 41}
	fmt.Println("The slice given here is:", numbers)
	
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	fmt.Println("The slice after reversal is:", numbers)
}