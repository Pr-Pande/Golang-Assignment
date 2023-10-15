package main

import (
	"fmt"
	"struct-func-pkgs/model"
	"struct-func-pkgs/shape"
)

func main() {
	p := model.Circle{
		Radius: 7,
	}
	
	res := shape.AreaOfCircle(p)
	fmt.Printf("Area of the circle of radius %g is %.2f", p.Radius, res)
}