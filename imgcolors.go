package imgcolors

import (
	"image"
	"image/color"
)

// NewImage create image.Image using colors.
func NewImage(colors []color.Color, width, height uint) image.Image {
	w := int(width)
	h := int(height)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	count := float64(len(colors))
	fw := float64(width)

	for x := 0; x < w; x++ {
		i := int(remap(float64(x), 0, fw, 0, count))
		for y := 0; y < h; y++ {
			img.Set(x, y, colors[i])
		}
	}
	return img
}

func remap(value, a, b, c, d float64) float64 {
	return (value-a)*((d-c)/(b-a)) + c
}
