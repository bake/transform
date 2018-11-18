package transform

import (
	"image"
	"image/draw"
	"math"
)

// Rotate by deg degree around its own center (m, n).
//  |x'|   |1 0 m|   |cos(d) -sin(d) 0|   |1 0 -m|
//  |y'| = |0 1 n| * |sin(d)  cos(d) 0| * |0 1 -n|
//  |z'|   |0 0 1|   |     0       0 1|   |0 0  1|
func Rotate(im draw.Image, deg float64) (draw.Image, error) {
	deg = deg * (math.Pi / 180)
	return exec(im, func(x, y int) (int, int) {
		x2 := math.Cos(deg)*float64(x) - math.Sin(deg)*float64(y)
		y2 := math.Sin(deg)*float64(x) + math.Cos(deg)*float64(y)
		return int(math.Round(x2)), int(math.Round(y2))
	})
}

// Scale by m and n.
//  |x'|   |m 0 0|   |x|
//  |y'| = |0 n 0| * |y|
//  |z'|   |0 0 1|   |1|
func Scale(im draw.Image, m, n float64) (draw.Image, error) {
	return exec(im, func(x, y int) (int, int) {
		return int(math.Round(float64(x) * m)), int(math.Round(float64(y) * n))
	})
}

// Shear by m and n.
//  |x'|   |1 n 0|   |x|
//  |y'| = |m 1 0| * |y|
//  |z'|   |0 0 1|   |1|
func Shear(im draw.Image, m, n float64) (draw.Image, error) {
	return exec(im, func(x, y int) (int, int) {
		return x + int(math.Round(float64(y)*n)), int(math.Round(float64(x)*m)) + y
	})
}

// Translate matrix.
//  |x'|   |1 0 m|   |x|
//  |y'| = |0 1 n| * |y|
//  |z'|   |0 0 1|   |1|
func Translate(im draw.Image, m, n int) (draw.Image, error) {
	return exec(im, func(x, y int) (int, int) {
		return x + m, y + n
	})
}

// exec the transformation on an draw.Image.
func exec(src draw.Image, f func(x, y int) (int, int)) (draw.Image, error) {
	b := src.Bounds()
	dst := image.NewRGBA(b)
	type M struct{ x1, y1, x2, y2 int }
	m, n := (b.Max.X-b.Min.X)/2, (b.Max.Y-b.Min.Y)/2
	c := make(chan M)
	defer close(c)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			go func(x, y int) {
				x2, y2 := f(x-m, y-n)
				c <- M{x, y, x2 + m, y2 + n}
			}(x, y)
		}
	}

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			m := <-c
			if m.x1 < b.Max.X && m.y1 < b.Max.Y && m.x2 < b.Max.X && m.y2 < b.Max.Y {
				dst.Set(m.x1, m.y1, src.At(m.x2, m.y2))
			}
		}
	}

	return dst, nil
}
