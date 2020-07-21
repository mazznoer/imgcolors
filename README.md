# imgcolors

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/imgcolors?tab=doc)

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
    colors := make([]color.Color, 17)
    for i := range colors {
        colors[i] = randColor()
    }
    img := imgcolors.Horizontal(colors, 900, 70)
    // ...
}

func rand255() uint8 {
    return uint8(rand.Intn(255))
}

func randColor() color.Color {
    return color.RGBA{rand255(), rand255(), rand255(), 255}
}
```
![example output](/examples/horizontal.png "Example output")
