package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Intn(10) + 1 //Generate a random number between 1 and 10
	
	var userNum int
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have chosen a random number between 1 and 10. Try to guess it!")

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&userNum)
		if userNum < 1 || userNum > 10 {
			fmt.Println("Please enter a number between 1 and 10")
		} else if userNum < num {
			fmt.Println("Too low")
		} else if userNum > num {
			fmt.Println("Too high")
		} else {
			fmt.Printf("You guessed it right")
			break
		}
	}
}
