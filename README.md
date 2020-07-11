# imgcolors

Create image from colors

## Example
```go
package main

import (
	"image/color"
	"math/rand"

	"github.com/mazznoer/imgcolors"
)

func main() {
	colors := make([]color.Color, 15)
	for i := range colors {
		colors[i] = color.RGBA{rand255(), rand255(), rand255(), 255}
	}
	img := imgcolors.NewImage(colors, 900, 90)
}

func rand255() uint8 {
	return uint8(rand.Intn(255))
}
```
![example output](/examples/random-colors.png "Example output")
