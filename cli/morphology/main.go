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
	out := flag.String("out", "out.gif", "Out file")
	flag.Parse()

	m := image.NewGray(image.Rect(0, 0, 11, 10))
	m.Pix = []uint8{
		000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
		000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
		000, 000, 255, 255, 000, 000, 000, 000, 000, 000, 000,
		000, 000, 255, 255, 000, 000, 255, 255, 255, 255, 000,
		000, 000, 255, 255, 000, 000, 255, 255, 255, 255, 000,
		000, 000, 255, 255, 255, 255, 255, 255, 255, 255, 000,
		000, 000, 255, 255, 255, 255, 255, 255, 255, 255, 000,
		000, 000, 000, 000, 000, 000, 255, 255, 255, 255, 000,
		000, 000, 000, 000, 000, 000, 255, 255, 255, 255, 000,
		000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
	}
	s := []uint8{
		000, 255, 000,
		255, 255, 255,
		000, 255, 000,
	}

	p, err := transform.Dilate(m.Pix, 11, s, 3)
	if err != nil {
		log.Fatal(err)
	}
	m.Pix = p

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gif.Encode(f, m, nil); err != nil {
		log.Fatal(err)
	}
}
