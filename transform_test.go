package transform

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	tt := []struct {
		pix, res []uint8
		width    int
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
			6,
			90,
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("rotate %d", i)
		t.Run(n, func(t *testing.T) {
			for j, p := range Rotate(tc.pix, tc.width, tc.deg) {
				if p != tc.res[j] {
					t.Fatalf("on %dx%d: expected %d got %d", j%11, j/11, tc.res[j], p)
				}
			}
		})
	}
}

func TestShear(t *testing.T) {
	tt := []struct {
		pix, res []uint8
		width    int
		m, n     float32
	}{
		{
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 255, 000,
				000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000,
				000, 255, 000, 000,
				000, 000, 000, 000,
				000, 000, 000, 255,
			},
			4,
			.5,
			.5,
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("shear %d", i)
		t.Run(n, func(t *testing.T) {
			for j, p := range Shear(tc.pix, tc.width, tc.m, tc.n) {
				if p != tc.res[j] {
					t.Fatalf("on %dx%d: expected %d got %d", j%11, j/11, tc.res[j], p)
				}
			}
		})
	}
}
