package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	concatenatedSlice := append(slice1, slice2...)

	fmt.Println(concatenatedSlice)
}
