package transform

import (
	"math"
	"sync"
)

// Translate matrix
//  |x'|   |1 0 m|   |x|
//  |y'| = |0 1 n| * |y|
//  |z'|   |0 0 1|   |1|
func Translate(p []uint8, pw, m, n int) []uint8 {
	return exec(p, pw, func(x, y int) (int, int, int) {
		return x + m, y + n, 1
	})
}

// Rotate by deg degree around its own center (m, n)
//  |x'|   |1 0 m|   |cos(d) -sin(d) 0|   |1 0 -m|
//  |y'| = |0 1 n| * |sin(d)  cos(d) 0| * |0 1 -n|
//  |z'|   |0 0 1|   |     0       0 1|   |0 0  1|
func Rotate(p []uint8, pw int, deg float64) []uint8 {
	deg = deg * (math.Pi / 180)
	m := pw / 2
	n := len(p) / pw / 2
	return exec(p, pw, func(x, y int) (int, int, int) {
		x0 := float64(x)
		y0 := float64(y)
		m0 := float64(m)
		n0 := float64(n)
		x1 := m0 + math.Cos(deg)*(x0-m0) - math.Sin(deg)*(y0-n0)
		y1 := n0 + math.Sin(deg)*(x0-m0) + math.Cos(deg)*(y0-n0)

		return int(x1), int(y1), 1
	})
}

// RotateC by deg degree around its own center (m, n)
//  |x'|   |1 0 m|   |cos(d) -sin(d) 0|   |1 0 -m|
//  |y'| = |0 1 n| * |sin(d)  cos(d) 0| * |0 1 -n|
//  |z'|   |0 0 1|   |     0       0 1|   |0 0  1|
func RotateC(p []uint8, pw int, deg float64) []uint8 {
	deg = deg * (math.Pi / 180)
	m := pw / 2
	n := len(p) / pw / 2
	return exec(p, pw, func(x, y int) (int, int, int) {
		x0 := float64(x)
		y0 := float64(y)
		m0 := float64(m)
		n0 := float64(n)
		x1 := m0 + math.Cos(deg)*(x0-m0) - math.Sin(deg)*(y0-n0)
		y1 := n0 + math.Sin(deg)*(x0-m0) + math.Cos(deg)*(y0-n0)

		return int(x1), int(y1), 1
	})
}

// Scale by m and n
//  |x'|   |m 0 0|   |x|
//  |y'| = |0 n 0| * |y|
//  |z'|   |0 0 1|   |1|
func Scale(p []uint8, pw int, m, n float64) []uint8 {
	return exec(p, pw, func(x, y int) (int, int, int) {
		return int(float64(x) * m), int(float64(y) * n), 1
	})
}

// Shear by m and n
//  |x'|   |1 n 0|   |x|
//  |y'| = |m 1 0| * |y|
//  |z'|   |0 0 1|   |1|
func Shear(p []uint8, pw int, m, n float32) []uint8 {
	return exec(p, pw, func(x, y int) (int, int, int) {
		return x + int(float32(y)*n), int(float32(x)*m) + y, 1
	})
}

func exec(p []uint8, pw int, f func(x, y int) (int, int, int)) []uint8 {
	type M struct{ i, j int }

	q := make([]uint8, len(p))
	c := make(chan M, len(p))
	defer close(c)

	for i := range p {
		go func(i int) {
			x, y, _ := f(i%pw, i/pw)
			c <- M{i, x + y*pw}
		}(i)
	}

	for range p {
		m := <-c
		if m.j >= 0 && m.j < len(p) {
			q[m.j] = p[m.i]
		}
	}

	return q
}

func execWG(p []uint8, pw int, f func(x, y int) (int, int, int)) []uint8 {
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
