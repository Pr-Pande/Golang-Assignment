package main

import "fmt"

type Rectangle struct {
	Width   int
	Height int
}

func (r Rectangle) calcArea() int {
	area := r.Width * r.Height
	return area
}

func (r Rectangle) calcPerimeter() int {
	perimeter := 2 * (r.Width + r.Height)
	return perimeter
}

func main() {
	rec := Rectangle{
		Width:   8,
		Height: 10,
	}

	fmt.Printf("Given dimensions of rectangle\nWidth: %d\nHeight: %d\n", rec.Width, rec.Height)

	fmt.Printf("Area of rectangle = %d\n", rec.calcArea())
	fmt.Printf("Perimeter of rectangle = %d", rec.calcPerimeter())
}
