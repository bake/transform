package main

import (
	"flag"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"

	"git.192k.pw/bake/transform"
)

func main() {
	in := flag.String("i", "in.gif", "Input file")
	out := flag.String("o", "out.gif", "Out file")
	w := flag.Int("width", 255, "Width")
	h := flag.Int("height", 255, "Height")
	b := flag.Int("border", 0, "Border")
	flag.Parse()

	r, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	img, err := gif.Decode(r)
	if err != nil {
		log.Fatal(err)
	}

	m := image.NewGray(image.Rect(0, 0, *w, *h))
	for i := range m.Pix {
		x := i % *w
		y := i / *w
		if x > *b && x < *w-*b && y > *b && y < *h-*b {
			m.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	m.Pix = transform.Shear(m.Pix, *w, .25, .25)
	m.Pix = transform.Scale(m.Pix, *w, .75, .5)
	m.Pix = transform.Rotate(m.Pix, *w, 45)
	m.Pix = transform.Translate(m.Pix, *w, *w/2, *w/8)

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := gif.Encode(f, m, nil); err != nil {
		log.Fatal(err)
	}
}
