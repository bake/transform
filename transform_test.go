package transform

import (
	"fmt"
	"image"
	"testing"
)

func TestRotate(t *testing.T) {
	im := image.NewGray(image.Rectangle{image.Point{X: 0, Y: 0}, image.Point{X: 6, Y: 6}})
	tt := []struct {
		pix, res []uint8
		deg      float64
	}{
		{
			[]uint8{
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 255, 255, 000, 000,
				000, 000, 000, 255, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 255, 000,
				000, 000, 000, 255, 255, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
			},
			90,
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("rotate %d", i)
		t.Run(n, func(t *testing.T) {
			im.Pix = tc.pix
			im, err := Rotate(im, tc.deg)
			if err != nil {
				t.Fatal(err)
			}
			for j, p := range im.(*image.Gray).Pix {
				if p != tc.res[j] {
					t.Fatalf("on %dx%d: expected %d got %d", j%11, j/11, tc.res[j], p)
				}
			}
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	im := image.NewGray(image.Rectangle{image.Point{X: 0, Y: 0}, image.Point{X: 6, Y: 6}})
	im.Pix = []uint8{
		000, 000, 000, 000, 000, 000,
		000, 000, 000, 000, 000, 000,
		000, 000, 255, 255, 000, 000,
		000, 000, 000, 255, 000, 000,
		000, 000, 000, 000, 000, 000,
		000, 000, 000, 000, 000, 000,
	}

	for n := 0; n < b.N; n++ {
		Rotate(im, 90)
	}
}

func TestShear(t *testing.T) {
	im := image.NewGray(image.Rectangle{image.Point{X: 0, Y: 0}, image.Point{X: 6, Y: 6}})
	tt := []struct {
		pix, res []uint8
		m, n     float32
	}{
		{
			[]uint8{
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 255, 255, 000, 000,
				000, 000, 000, 255, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 000, 000,
				000, 000, 000, 255, 255, 000,
				000, 000, 000, 000, 255, 000,
				000, 000, 000, 000, 000, 000,
			},
			.5,
			.5,
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("shear %d", i)
		t.Run(n, func(t *testing.T) {
			im.Pix = tc.pix
			im, err := Shear(im, tc.m, tc.n)
			if err != nil {
				t.Fatal(err)
			}
			for j, p := range im.(*image.Gray).Pix {
				if p != tc.res[j] {
					t.Fatalf("on %dx%d: expected %d got %d", j%11, j/11, tc.res[j], p)
				}
			}
		})
	}
}
