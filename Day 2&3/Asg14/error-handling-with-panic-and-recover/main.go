package main

import (
	"errors"
	"fmt"
)

func divide(numerator, denominator int) int {
	if denominator == 0 {
		panic("division by zero is not allowed")
	}
	return numerator / denominator
}

func safeDivide(numerator, denominator int) (int, error) {
	var result int
	var err error

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recovered: %v\n", r)
			err = errors.New("division by zero")
		}
	}()

	result = divide(numerator, denominator)
	return result, err
}

func main() {
	numerator1, denominator1 := 10, 2
	numerator2, denominator2 := 8, 0

	result1, err1 := safeDivide(numerator1, denominator1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Result 1:", result1)
	}

	result2, err2 := safeDivide(numerator2, denominator2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result 2:", result2)
	}
}
