package main

import (
	"fmt"
	"project-temperature/temperature"
)

func main() {
	var temp float64
	fmt.Println("Enter temperature in Fahrenheit:")
	fmt.Scanln(&temp)
	fmt.Printf("Temperature in Celsius: %g", temperature.TempConvFromFToC(temp))
}
