package main

import "fmt"

func main() {
	numbers := []int{75, 42, 63, 8, 41}
	
	val := largestElement(numbers)

    fmt.Println("The largest element in the slice is:")
    fmt.Println(val)
}

func largestElement(slice []int) int {
	var max int
	max = slice[0]  
	for _, ele := range slice {
	   if ele > max { 
		  max = ele  
	   }
	}
	return max
}