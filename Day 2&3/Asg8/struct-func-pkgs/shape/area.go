package shape

import (
	"math"
	"struct-func-pkgs/model"
)

var pi = math.Pi

func AreaOfCircle(c model.Circle) float64 {
	area := pi * c.Radius * c.Radius
	return area
}