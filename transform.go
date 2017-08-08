package transform

import "math"
import "sync"

// Translate matrix
func Translate(p []uint8, pw, m, n int) []uint8 {
	return exec(p, pw, func(x, y int) (int, int, int) {
		return translate(x, y, m, n)
	})
}

func translate(x, y, m, n int) (int, int, int) {
	return x + m, y + n, 1
}

// Rotate by deg degree
func Rotate(p []uint8, pw int, deg float64) []uint8 {
	deg = deg * (math.Pi / 180)
	return exec(p, pw, func(x, y int) (int, int, int) {
		return rotate(x, y, deg)
	})
}

func rotate(x, y int, deg float64) (int, int, int) {
	x1 := float64(x)*math.Cos(deg) - float64(y)*math.Sin(deg)
	y1 := float64(x)*math.Sin(deg) + float64(y)*math.Cos(deg)
	return int(x1), int(y1), 1
}

// Scale by m and n
func Scale(p []uint8, pw int, m, n float64) []uint8 {
	return exec(p, pw, func(x, y int) (int, int, int) {
		return scale(x, y, m, n)
	})
}

func scale(x, y int, m, n float64) (int, int, int) {
	return int(float64(x) * m), int(float64(y) * n), 1
}

// Shear by m and n
func Shear(p []uint8, pw int, m, n float32) []uint8 {
	return exec(p, pw, func(x, y int) (int, int, int) {
		return shear(x, y, m, n)
	})
}

func shear(x, y int, m, n float32) (int, int, int) {
	return x + int(float32(y)*n), int(float32(x)*m) + y, 1
}

func exec(p []uint8, pw int, f func(x, y int) (int, int, int)) []uint8 {
	q := make([]uint8, len(p))
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(p))
	for i := range p {
		go func(i int) {
			x, y, _ := f(i%pw, i/pw)
			j := x + y*pw
			if j >= 0 && j < len(p) {
				mu.Lock()
				q[j] = p[i]
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return q
}
