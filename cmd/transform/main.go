package main

import (
	"flag"
	"image/gif"
	"log"
	"os"

	"git.192k.pw/bake/transform"
)

func main() {
	in := flag.String("i", "in.gif", "Input file")
	out := flag.String("o", "out.gif", "Out file")
	degree := flag.Float64("deg", 90, "Degree")
	flag.Parse()

	r, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	img, err := gif.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	if img, err = transform.Rotate(img, *degree); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := gif.Encode(f, img, nil); err != nil {
		log.Fatal(err)
	}
}
