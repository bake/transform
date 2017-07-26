package transform

import (
	"fmt"
	"testing"
)

func TestDilate(t *testing.T) {
	tt := []struct {
		pix, s, res []uint8
	}{
		{
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 255, 000,
				000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000,
				000, 255, 255,
				000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 255,
				000, 000, 255, 255,
				000, 000, 000, 000,
			},
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("dilate %d", i)
		t.Run(n, func(t *testing.T) {
			for i, p := range Dilate(tc.pix, 4, tc.s, 3) {
				if p != tc.res[i] {
					t.Fatalf("on %dx%d: expected %d got %d", i%11, i/11, tc.res[i], p)
				}
			}
		})
	}
}

func TestErode(t *testing.T) {
	tt := []struct {
		pix, s, res []uint8
	}{
		{
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 255, 000,
				000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000,
				000, 255, 255,
				000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000,
				000, 255, 000, 000,
				000, 000, 000, 000,
				000, 000, 000, 000,
			},
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("erode %d", i)
		t.Run(n, func(t *testing.T) {
			for i, p := range Erode(tc.pix, 4, tc.s, 3) {
				if p != tc.res[i] {
					t.Fatalf("on %dx%d: expected %d got %d", i%11, i/11, tc.res[i], p)
				}
			}
		})
	}
}

func TestOpen(t *testing.T) {
	tt := []struct {
		pix, s, res []uint8
	}{
		{
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 255, 000,
				000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000,
				000, 255, 255,
				000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 000, 000,
				000, 000, 000, 000,
			},
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("open %d", i)
		t.Run(n, func(t *testing.T) {
			for i, p := range Open(tc.pix, 4, tc.s, 3) {
				if p != tc.res[i] {
					t.Fatalf("on %dx%d: expected %d got %d", i%11, i/11, tc.res[i], p)
				}
			}
		})
	}
}

func TestClose(t *testing.T) {
	tt := []struct {
		pix, s, res []uint8
	}{
		{
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 255, 000,
				000, 000, 000, 000,
			},
			[]uint8{
				000, 000, 000,
				000, 255, 255,
				000, 000, 000,
			},
			[]uint8{
				000, 000, 000, 000,
				000, 255, 255, 000,
				000, 000, 255, 000,
				000, 000, 000, 000,
			},
		},
	}

	for i, tc := range tt {
		n := fmt.Sprintf("close %d", i)
		t.Run(n, func(t *testing.T) {
			for i, p := range Close(tc.pix, 4, tc.s, 3) {
				if p != tc.res[i] {
					t.Fatalf("on %dx%d: expected %d got %d", i%11, i/11, tc.res[i], p)
				}
			}
		})
	}
}
