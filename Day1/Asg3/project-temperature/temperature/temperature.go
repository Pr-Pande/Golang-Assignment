package temperature

func TempConvFromFToC(f float64) float64 {
	c := (f - 32) * 5 / 9
	return c
}