// +build ignore

package main

import (
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/mazznoer/imgcolors"
)

func main() {
	img := imgcolors.Horizontal(randomColors(17), 900, 70)

	file, err := os.Create("colors.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	png.Encode(file, img)
}

func randomColors(n uint) []color.Color {
	colors := make([]color.Color, n)
	for i := range colors {
		colors[i] = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	}
	return colors
}
