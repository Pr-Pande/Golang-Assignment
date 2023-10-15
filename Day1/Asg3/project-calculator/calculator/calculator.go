package calculator

func Square(x int) int {
	return x*x
}

func Sum(x, y int) int {
	return x + y
}

func Diff(x, y int) int {
	return x - y
}

func Prod(x, y int) int {
	return x*y
}

func QuoAndRem(x, y int) (float64, int) {
	quo := float64(x / y)
	rem := x % y
	return quo, rem
}


