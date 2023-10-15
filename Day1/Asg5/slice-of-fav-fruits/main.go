package main

import "fmt"

func main() {
	fruit := []string{"Mango", "Apple", "Papaya", "Banana"}

	for _, f := range fruit {
		fmt.Println(f)
	}
}