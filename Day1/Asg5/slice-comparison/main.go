package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{10, 20, 30, 40, 50}
	fmt.Println("The elements of slice1 are:", myslice1)
	myslice2 := []int{10, 20, 30, 40, 50}
	fmt.Println("The elements of slice2 are:", myslice2)

	if len(myslice1) != len(myslice2) { //if condition is not satisfied print false
		fmt.Println("Two slices are not equal")
		return
	}

	for i, element := range myslice1 { // use for loop to check equality
		if element != myslice2[i] {
			fmt.Println("Two slices are not equal")
		}
	}
	fmt.Println("Two slices are equal")

}
