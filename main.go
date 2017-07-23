package main

import (
	"flag"
	"image"
	"image/gif"
	"log"
	"math"
	"os"
)

func main() {
	out := flag.String("out", "out.gif", "Out file")
	deg := flag.Float64("deg", 0, "Degree")
	flag.Parse()

	w := 255
	h := 255

	m := image.NewGray(image.Rect(0, 0, w, h))
	n := image.NewGray(image.Rect(0, 0, w, h))

	for i := 0; i < len(m.Pix); i++ {
		m.Pix[i] = 255 // uint8(i % w)
	}

	for i := 0; i < len(m.Pix); i++ {
		x := i % w
		y := i / w
		if x < 50 || x > w-50 || y < 50 || y > h-50 {
			continue
		}

		x, y, _ = rotate(x, y, *deg)
		// x, y, _ = translate(x, y, x%16, y%16)
		j := x + y*w
		if j >= 0 && j < len(n.Pix) {
			n.Pix[j] = m.Pix[i]
		}
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gif.Encode(f, n, nil); err != nil {
		log.Fatal(err)
	}
}

func translate(x, y, m, n int) (int, int, int) {
	return x + m, y + n, 1
}

func rotate(x, y int, deg float64) (int, int, int) {
	deg = deg * (math.Pi / 180)
	x1 := float64(x)*math.Cos(deg) - float64(y)*math.Sin(deg)
	y1 := float64(x)*math.Sin(deg) + float64(y)*math.Cos(deg)
	z1 := float64(1)
	return int(x1), int(y1), int(z1)
}

func scale(x, y, m, n int) (int, int, int) {
	return x * m, y * n, 1
}
