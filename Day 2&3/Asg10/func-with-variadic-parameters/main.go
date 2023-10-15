package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, n := range nums {
        total += n
    }
    fmt.Println("\nSum:", total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    numSlice := []int{1, 2, 3, 4}
    sum(numSlice...)
}