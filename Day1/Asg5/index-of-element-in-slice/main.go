package main

import "fmt"

func main() {

	numbers := []int{75, 42, 63, 8, 41, 25}

	fmt.Println("The slice given here is:", numbers)

	// Call the search function and store the value in a variable named val
	res := searchNum(numbers, 8)
	fmt.Println("The value to be searched from the slice is:", 8)

	if res != -1 {
		fmt.Println("The element is found in slice at index:", res)
	} else {
		fmt.Println("The element was not found in the slice")
	}
}

func searchNum(slice []int, key int) int {
	for i, element := range slice {
		if element == key { // check the condition if its true return index
			return i
		}
	}
	return -1
}
