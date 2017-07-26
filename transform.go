package transform

import "math"

// Translate matrix
func Translate(p []uint8, pw, m, n int) []uint8 {
	q := make([]uint8, len(p))

	for i := range p {
		x := i % pw
		y := i / pw

		x, y, _ = scale(x, y, m, n)
		j := x + y*pw
		if j >= 0 && j < len(q) {
			q[j] = p[i]
		}
	}

	return q
}

func translate(x, y, m, n int) (int, int, int) {
	return x + m, y + n, 1
}

// Rotate by deg degree
func Rotate(p []uint8, pw int, deg float64) []uint8 {
	deg = deg * (math.Pi / 180)
	q := make([]uint8, len(p))

	for i := range p {
		x, y, _ := rotate(i%pw, i/pw, deg)
		j := x + y*pw
		if j >= 0 && j < len(q) {
			q[j] = p[i]
		}
	}

	return q
}

func rotate(x, y int, deg float64) (int, int, int) {
	x1 := float64(x)*math.Cos(deg) - float64(y)*math.Sin(deg)
	y1 := float64(x)*math.Sin(deg) + float64(y)*math.Cos(deg)
	return int(x1), int(y1), 1
}

// Scale by m and n
func Scale(p []uint8, pw, m, n int) []uint8 {
	q := make([]uint8, len(p))

	for i := range p {
		x := i % pw
		y := i / pw

		x, y, _ = scale(x, y, m, n)
		j := x + y*pw
		if j >= 0 && j < len(q) {
			q[j] = p[i]
		}
	}

	return q
}

func scale(x, y, m, n int) (int, int, int) {
	return x * m, y * n, 1
}
