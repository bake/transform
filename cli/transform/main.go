package main

import (
	"flag"
	"image"
	"image/gif"
	"log"
	"os"

	"git.192k.pw/bake/transform"
)

func main() {
	out := flag.String("o", "out.gif", "Out file")
	w := flag.Int("width", 255, "Width")
	h := flag.Int("height", 255, "Height")
	b := flag.Int("border", 50, "Border")
	deg := flag.Float64("deg", 0, "Degree")
	flag.Parse()

	m := image.NewGray(image.Rect(0, 0, *w, *h))
	for i := range m.Pix {
		x := i % *w
		y := i / *w
		if x > *b && x < *w-*b && y > *b && y < *h-*b {
			m.Pix[i] = 255
		}
	}
	m.Pix = transform.Rotate(m.Pix, *w, *deg)

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := gif.Encode(f, m, nil); err != nil {
		log.Fatal(err)
	}
}
