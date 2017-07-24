package transform

import "fmt"

// Dilate pixels p with mask s
func Dilate(p []uint8, pw int, s []uint8, sw int) ([]uint8, error) {
	if sw%2 == 0 {
		return nil, fmt.Errorf("the mask has to be odd numbered")
	}

	ph := len(p) / pw
	sh := len(s) / sw
	r := make([]uint8, len(p))

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == 0 {
				continue
			}

			x := i%pw + j%sw - sw/2
			y := i/pw + j/sh - sh/2
			if x < 0 || y < 0 || x >= pw || y >= ph {
				continue
			}
			if s[j] > p[x+y*pw] {
				continue
			}
			r[i] = s[j]
		}
	}

	return r, nil
}

// Erode pixels p with mask s
func Erode(p []uint8, pw int, s []uint8, sw int) ([]uint8, error) {
	if sw%2 == 0 {
		return nil, fmt.Errorf("the mask has to be odd numbered")
	}

	ph := len(p) / pw
	sh := len(s) / sw
	r := make([]uint8, len(p))

	for i := 0; i < len(p); i++ {
		ok := true
		for j := 0; j < len(s); j++ {
			if s[j] == 0 {
				continue
			}

			x := i%pw + j%sw - sw/2
			y := i/pw + j/sh - sh/2
			if x < 0 || y < 0 || x >= pw || y >= ph {
				ok = false
				break
			}
			if s[j] > p[x+y*pw] {
				ok = false
				break
			}
		}
		if ok {
			r[i] = p[i]
		}
	}

	return r, nil
}

func Open(p []uint8, pw int, s []uint8, sw int) ([]uint8, error) {
	p, err := Erode(p, pw, s, sw)
	if err != nil {
		return nil, err
	}
	return Dilate(p, pw, s, sw)
}

func Close(p []uint8, pw int, s []uint8, sw int) ([]uint8, error) {
	p, err := Dilate(p, pw, s, sw)
	if err != nil {
		return nil, err
	}
	return Erode(p, pw, s, sw)
}
