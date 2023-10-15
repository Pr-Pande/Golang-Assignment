package main

import "fmt"

func main() {
	fruits := make(map[string]int)
	fruits["Grapes"] = 10
	fruits["Mango"] = 20
	fruits["Banana"] = 30
	fruits["Apple"] = 40

	for k, v := range fruits {
		fmt.Printf("Fruit = %s, Quantity = %d\n", k, v)
	}

}
