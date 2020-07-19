package image

import (
	"image"
	"image/color"
)

type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}
