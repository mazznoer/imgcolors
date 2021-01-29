# imgcolors

[![Release](https://img.shields.io/github/release/mazznoer/imgcolors.svg)](https://github.com/mazznoer/imgcolors/releases/latest)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/imgcolors?tab=doc)
[![Build Status](https://travis-ci.org/mazznoer/imgcolors.svg?branch=master)](https://travis-ci.org/mazznoer/imgcolors)
[![Build Status](https://github.com/mazznoer/imgcolors/workflows/Go/badge.svg)](https://github.com/mazznoer/imgcolors/actions)
[![go report](https://goreportcard.com/badge/github.com/mazznoer/imgcolors)](https://goreportcard.com/report/github.com/mazznoer/imgcolors)
[![codecov](https://codecov.io/gh/mazznoer/imgcolors/branch/master/graph/badge.svg)](https://codecov.io/gh/mazznoer/imgcolors)

Create image from colors

## Examples

### Horizontal
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
![example output](docs/images/horizontal.png)

[Try it online](https://play.golang.org/p/7zaL_OQ4Gbf)

### Vertical
```go
colors := make([]color.Color, 10)

for i := range colors {
    colors[i] = randColor()
}
img := imgcolors.Vertical(colors, 45, 45*10)
```
![example output](docs/images/vertical.png "Example output")

### Square
```go
import "image/color/palette"

img := imgcolors.Square(palette.Plan9, 900)
```
![example output](docs/images/square.png "Example output")

### Conic
```go
colors := make([]color.Color, 17)

for i := range colors {
    colors[i] = randColor()
}
img := imgcolors.Conic(colors, 800, 800)
```
![example output](docs/images/conic.png "Example output")

### Radial
```go
colors := make([]color.Color, 13)

for i := range colors {
    colors[i] = randColor()
}
img := imgcolors.Radial(colors, 800, 800)
```
![example output](docs/images/radial.png "Example output")
