package imgcolors

import (
	"image/color"
	"math"
	"math/rand"
	"testing"
)

func TestHorizontal(t *testing.T) {
	colors := randColors(17)
	img := Horizontal(colors, 800, 50)
	for i, v := range linspace(0, 800-1, 17) {
		c := img.At(int(v), 25)
		if c != colors[i] {
			t.Errorf("%v != %v", c, colors[i])
		}
	}
}

func TestVertical(t *testing.T) {
	colors := randColors(35)
	img := Vertical(colors, 35, 400)
	for i, v := range linspace(0, 400-1, 35) {
		c := img.At(10, int(v))
		if c != colors[i] {
			t.Errorf("%v != %v", c, colors[i])
		}
	}
}

func TestSquare(t *testing.T) {
	colors := randColors(57)
	img := Square(colors, 350)
	n := uint(math.Ceil(math.Sqrt(float64(len(colors)))))
	i := 0
	for _, y := range linspace(0, 350-1, n) {
		for _, x := range linspace(0, 350-1, n) {
			if i >= 57 {
				break
			}
			c := img.At(int(x), int(y))
			if c != colors[i] {
				t.Errorf("%v != %v", c, colors[i])
			}
			i++
		}
	}
}

func TestConic(t *testing.T) {
	colors := randColors(11)
	img := Conic(colors, 500, 500)
	cx := 250.0
	cy := 250.0
	p := math.Pi * 2 / 11
	s := math.Pi + p/2
	for i := 0; i < 11; i++ {
		x := cx + 100*math.Cos(s+p*float64(i))
		y := cy + 100*math.Sin(s+p*float64(i))
		c := img.At(int(x), int(y))
		if c != colors[i] {
			t.Errorf("%v != %v", c, colors[i])
		}
	}
}

func TestRadial(t *testing.T) {
	colors := randColors(11)
	img := Radial(colors, 500, 500)
	for i, v := range linspace(0, 250-1, 11) {
		c := img.At(250, 250+int(v))
		if c != colors[i] {
			t.Errorf("%v != %v", c, colors[i])
		}
	}
}

func randColors(n uint) []color.Color {
	colors := make([]color.Color, n)
	for i := range colors {
		colors[i] = color.RGBA{rand255(), rand255(), rand255(), 255}
	}
	return colors
}

func rand255() uint8 {
	return uint8(rand.Intn(255))
}

func linspace(min, max float64, n uint) []float64 {
	d := max - min
	l := float64(n - 1)
	res := make([]float64, n)
	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}
