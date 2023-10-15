package main

import "fmt"

func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var n int
	fmt.Println("Enter a number:")
	fmt.Scanln(&n)
	
    result := factorial(n)
    fmt.Printf("The factorial of %d is %d\n", n, result)
}