package tour

import (
	"image"
	"image/color"
)

// Image has height and width
type Image struct {
	h int
	w int
}

// ColorModel returns always RGBA64Model.
func (i *Image) ColorModel() color.Model {
	return color.RGBA64Model
}

// Bounds creates Rectangle.
func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At returns color at the point.
// If sum of x and y is even, returns black, else returns white.
func (i *Image) At(x, y int) color.Color {
	if x+y%2 == 0 {
		return color.Black
	}
	return color.White
}

// Pic creates Image
// height is dy.
// width is dx.
func Pic(dx, dy int) *Image {
	return &Image{dy, dx}
}
