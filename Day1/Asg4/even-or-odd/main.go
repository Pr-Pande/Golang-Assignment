package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)
	checkEvenOdd(num)
}

func checkEvenOdd(a int) {
	if a & 1 == 0 {
        fmt.Println("Number is Even.\n")
    } else {
        fmt.Printf("Number is Odd.\n")
    }
}