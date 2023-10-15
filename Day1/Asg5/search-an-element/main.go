package main

import "fmt"

func main() {
	numbers := []int{22, 44, 11, 77, 88}
	var num int

	fmt.Println("The slice given here is:", numbers)
	fmt.Println("\nEnter the number you want to search in slice:")
	fmt.Scanln(&num)

	val := searchEle(numbers, num)

	if val != -1 {
		fmt.Println("The element is found in slice at index:", val)
	} else {
		fmt.Println("The element was not found in the slice")
	}
}

func searchEle(slice []int, key int) int {
	for i, ele := range slice {
		if ele == key {
			return i
		}
	}
	return -1
}
