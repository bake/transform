package transform_test

import (
	"fmt"
	"image"
	"testing"

	"git.192k.pw/bake/transform"
)

func TestRotate(t *testing.T) {
	tt := []struct {
		w, h     int
		pix, res []uint8
		deg      float64
	}{
		{
			3, 3,
			[]uint8{
				000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 255, 255, 255, 255, 255, 255, 255, 255,
				000, 000, 000, 000, 000, 000, 000, 000, 255, 255, 255, 255,
			},
			[]uint8{
				000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000, 000,
				000, 000, 000, 000, 255, 255, 255, 255, 000, 000, 000, 000,
				255, 255, 255, 255, 255, 255, 255, 255, 000, 000, 000, 000,
			},
			90,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("rotate %d", i), func(t *testing.T) {
			src := image.NewRGBA(image.Rect(0, 0, tc.w, tc.h))
			src.Pix = tc.pix
			dst, err := transform.Rotate(src, tc.deg)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(dst.(*image.RGBA).Pix)
			for j, p := range dst.(*image.RGBA).Pix {
				if p != tc.res[j] {
					t.Fatalf("on %dx%d: expected %d, got %d", j%11, j/11, tc.res[j], p)
				}
			}
		})
	}
}
