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

	w := 11
	h := 8
	m := image.NewGray(image.Rect(0, 0, w, h))
	b := []uint8{
		000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
		000, 000, 255, 255, 000, 000, 000, 000, 255, 255, 000,
		000, 255, 255, 255, 000, 000, 255, 000, 255, 000, 000,
		000, 255, 255, 255, 000, 000, 255, 000, 000, 000, 000,
		000, 255, 255, 255, 255, 000, 255, 000, 000, 255, 000,
		000, 255, 255, 000, 255, 255, 000, 000, 255, 255, 000,
		000, 255, 255, 255, 255, 000, 000, 000, 255, 255, 000,
		000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
	}

	s := []uint8{
		000, 255, 255,
		000, 255, 000,
		000, 000, 000,
	}

	m.Pix = transform.Open(b, w, s, 3)

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := gif.Encode(f, m, nil); err != nil {
		log.Fatal(err)
	}
}
