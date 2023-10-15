package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 12}
	sum, avg := sumAndAverage(numbers)
	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", avg)
}

func sumAndAverage(slice []int) (int, int) {
	s := 0
	for _, i := range slice {
		s += i
	}

	a := s / len(slice)
	return s, a
}