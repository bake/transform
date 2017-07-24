package transform

import "math"

// Translate matrix
func Translate(x, y, m, n int) (int, int, int) {
	return x + m, y + n, 1
}

// Rotate by deg degree
func Rotate(x, y int, deg float64) (int, int, int) {
	deg = deg * (math.Pi / 180)
	x1 := float64(x)*math.Cos(deg) - float64(y)*math.Sin(deg)
	y1 := float64(x)*math.Sin(deg) + float64(y)*math.Cos(deg)
	z1 := float64(1)
	return int(x1), int(y1), int(z1)
}

// Scale by m and n
func Scale(x, y, m, n int) (int, int, int) {
	return x * m, y * n, 1
}
