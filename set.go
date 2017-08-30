package transform

import (
	"fmt"
	"image"
	"image/color"
)

// getSet returns a image.Image-Set-function and a new image.Image with the
// same dimensions and palette as the input image. Due to the lack of this
// function only NYCbCrA and YCbCr are unsupported.
func getSet(im image.Image) (func(int, int, color.Color), image.Image, error) {
	set := func(x, y int, c color.Color) {}
	switch tmp := im.(type) {
	case *image.Alpha:
		im = image.NewAlpha(tmp.Rect)
		set = im.(*image.Alpha).Set
	case *image.Alpha16:
		im = image.NewAlpha16(tmp.Rect)
		set = im.(*image.Alpha16).Set
	case *image.CMYK:
		im = image.NewCMYK(tmp.Rect)
		set = im.(*image.CMYK).Set
	case *image.Gray:
		im = image.NewGray(tmp.Rect)
		set = im.(*image.Gray).Set
	case *image.Gray16:
		im = image.NewGray16(tmp.Rect)
		set = im.(*image.Gray16).Set
	case *image.NRGBA:
		im = image.NewNRGBA(tmp.Rect)
		set = im.(*image.NRGBA).Set
	case *image.NRGBA64:
		im = image.NewNRGBA64(tmp.Rect)
		set = im.(*image.NRGBA64).Set
	case *image.Paletted:
		im = image.NewPaletted(tmp.Rect, tmp.Palette)
		set = im.(*image.Paletted).Set
	case *image.RGBA:
		im = image.NewRGBA(tmp.Rect)
		set = im.(*image.RGBA).Set
	case *image.RGBA64:
		im = image.NewRGBA64(tmp.Rect)
		set = im.(*image.RGBA64).Set
	default:
		return nil, nil, fmt.Errorf("%T has no function Set", im)
	}
	return set, im, nil
}
