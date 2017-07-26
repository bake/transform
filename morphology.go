package transform

// Dilate pixels p with mask s
func Dilate(p []uint8, pw int, s []uint8, sw int) []uint8 {
	ph := len(p) / pw
	sh := len(s) / sw
	r := make([]uint8, len(p))

	for i := 0; i < len(p); i++ {
		if p[i] != s[len(s)/2] {
			continue
		}

		for j := 0; j < len(s); j++ {
			if s[j] == 0 {
				continue
			}

			x := i%pw + j%sw - sw/2
			y := i/pw + j/sh - sh/2
			if x < 0 || y < 0 || x >= pw || y >= ph {
				continue
			}
			r[x+y*pw] = s[j]
		}
	}

	return r
}

// Erode pixels p with mask s
func Erode(p []uint8, pw int, s []uint8, sw int) []uint8 {
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

	return r
}

// Open executes erode and dilate
func Open(p []uint8, pw int, s []uint8, sw int) []uint8 {
	return Dilate(Erode(p, pw, s, sw), pw, s, sw)
}

// Close executes dilate and erode
func Close(p []uint8, pw int, s []uint8, sw int) []uint8 {
	return Erode(Dilate(p, pw, s, sw), pw, s, sw)
}
