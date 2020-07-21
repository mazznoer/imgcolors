package imgcolors

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

func Horizontal(colors []color.Color, width, height uint) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	fw := float64(width)
	n := len(colors)
	fn := float64(n)
	px := 0

	for x := 0; x < n; x++ {
		nx := int(remap(float64(x+1), 0, fn, 0, fw))
		draw.Draw(img, image.Rect(px, 0, nx, int(height)), &image.Uniform{colors[x]}, image.Point{}, draw.Src)
		px = nx
	}
	return img
}

func Vertical(colors []color.Color, width, height uint) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	fh := float64(height)
	n := len(colors)
	fn := float64(n)
	py := 0

	for y := 0; y < n; y++ {
		ny := int(remap(float64(y+1), 0, fn, 0, fh))
		draw.Draw(img, image.Rect(0, py, int(width), ny), &image.Uniform{colors[y]}, image.Point{}, draw.Src)
		py = ny
	}
	return img
}

func Square(colors []color.Color, size uint) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	n := int(math.Ceil(math.Sqrt(float64(len(colors)))))
	fsize := float64(size)
	fn := float64(n)
	px := 0
	py := 0

	for y := 0; y < n; y++ {
		ny := int(remap(float64(y+1), 0, fn, 0, fsize))

		for x := 0; x < n; x++ {
			i := y*n + x
			if i >= len(colors) {
				break
			}
			nx := int(remap(float64(x+1), 0, fn, 0, fsize))
			draw.Draw(img, image.Rect(px, py, nx, ny), &image.Uniform{colors[i]}, image.Point{}, draw.Src)
			px = nx
		}
		px = 0
		py = ny
	}
	return img
}

func Conic(colors []color.Color, width, height uint) image.Image {
	w := int(width)
	h := int(height)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	cx := w / 2
	cy := h / 2
	n := len(colors)
	fn := float64(len(colors))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := math.Atan2(float64(y-cy), float64(x-cx))
			i := int(math.Floor(remap(a, -math.Pi, math.Pi, 0, fn)))
			if i >= n {
				i = 0
			}
			img.Set(x, y, colors[i])
		}
	}
	return img
}

func remap(value, a, b, c, d float64) float64 {
	return (value-a)*((d-c)/(b-a)) + c
}
