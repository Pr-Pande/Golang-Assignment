package main

import "fmt"

func removeDuplicates(nums []int) []int {
	result := []int{}

	for _, num := range nums {
		if !contains(result, num) {
			result = append(result, num)
		}
	}

	return result
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {

	numbers := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println("The initial slice is:", numbers)

	// Remove duplicates from the numbers slice
	deduplicatedSlice := removeDuplicates(numbers)

	// Print the deduplicated slice
	fmt.Println("The deduplicated slice is:", deduplicatedSlice)
}
