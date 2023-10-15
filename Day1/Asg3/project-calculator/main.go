package main

import (
	"fmt"
	"project-calculator/calculator"
)

func main() {
	fmt.Println(calculator.Square(12))
	fmt.Println(calculator.Sum(650, 40))
	fmt.Println(calculator.Diff(6544, 3278))
	fmt.Println(calculator.Prod(24, 25))
	fmt.Println(calculator.QuoAndRem(63, 5))
}