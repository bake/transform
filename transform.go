package transform

import (
	"image"
	"math"
)

// Rotate by deg degree around its own center (m, n)
//  |x'|   |1 0 m|   |cos(d) -sin(d) 0|   |1 0 -m|
//  |y'| = |0 1 n| * |sin(d)  cos(d) 0| * |0 1 -n|
//  |z'|   |0 0 1|   |     0       0 1|   |0 0  1|
func Rotate(im image.Image, deg float64) (image.Image, error) {
	deg = deg * (math.Pi / 180)
	b := im.Bounds()
	m := (b.Max.X - b.Min.X) / 2
	n := (b.Max.Y - b.Min.Y) / 2
	return exec(im, func(x, y int) (int, int) {
		x0, y0 := float64(x), float64(y)
		m0, n0 := float64(m), float64(n)
		x1 := m0 + math.Cos(deg)*(x0-m0) - math.Sin(deg)*(y0-n0)
		y1 := n0 + math.Sin(deg)*(x0-m0) + math.Cos(deg)*(y0-n0)

		return int(x1), int(y1)
	})
}

// Scale by m and n
//  |x'|   |m 0 0|   |x|
//  |y'| = |0 n 0| * |y|
//  |z'|   |0 0 1|   |1|
func Scale(im image.Image, m, n float64) (image.Image, error) {
	return exec(im, func(x, y int) (int, int) {
		return int(float64(x) * m), int(float64(y) * n)
	})
}

// Shear by m and n
//  |x'|   |1 n 0|   |x|
//  |y'| = |m 1 0| * |y|
//  |z'|   |0 0 1|   |1|
func Shear(im image.Image, m, n float32) (image.Image, error) {
	return exec(im, func(x, y int) (int, int) {
		return x + int(float32(y)*n), int(float32(x)*m) + y
	})
}

// Translate matrix
//  |x'|   |1 0 m|   |x|
//  |y'| = |0 1 n| * |y|
//  |z'|   |0 0 1|   |1|
func Translate(im image.Image, m, n int) (image.Image, error) {
	return exec(im, func(x, y int) (int, int) {
		return x + m, y + n
	})
}

// exec the transformation on an image.Image.
func exec(im image.Image, f func(x, y int) (int, int)) (image.Image, error) {
	set, im2, err := getSet(im)
	if err != nil {
		return nil, err
	}

	type M struct{ x1, y1, x2, y2 int }
	b := im.Bounds()
	p := (b.Max.X - b.Min.X) * (b.Max.Y - b.Min.Y)
	c := make(chan M)
	defer close(c)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			go func(x, y int) {
				x2, y2 := f(x, y)
				c <- M{x, y, x2, y2}
			}(x, y)
		}
	}

	for i := 0; i < p; i++ {
		m := <-c
		if m.x1 < b.Max.X && m.y1 < b.Max.Y && m.x2 < b.Max.X && m.y2 < b.Max.Y {
			set(m.x2, m.y2, im.At(m.x1, m.y1))
		}
	}

	return im2, nil
}
