package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    cir := Circle{Radius: 5.0}
    rect := Rectangle{Width: 6.0, Height: 4.0}

    fmt.Printf("Circle Area: %.2f\n", cir.Area())
    fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
}