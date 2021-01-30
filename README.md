# imgcolors

[![Release](https://img.shields.io/github/release/mazznoer/imgcolors.svg)](https://github.com/mazznoer/imgcolors/releases/latest)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/imgcolors?tab=doc)
[![Build Status](https://travis-ci.org/mazznoer/imgcolors.svg?branch=master)](https://travis-ci.org/mazznoer/imgcolors)
[![Build Status](https://github.com/mazznoer/imgcolors/workflows/Go/badge.svg)](https://github.com/mazznoer/imgcolors/actions)
[![go report](https://goreportcard.com/badge/github.com/mazznoer/imgcolors)](https://goreportcard.com/report/github.com/mazznoer/imgcolors)
[![codecov](https://codecov.io/gh/mazznoer/imgcolors/branch/master/graph/badge.svg)](https://codecov.io/gh/mazznoer/imgcolors)

Create image from colors

```go
import "github.com/mazznoer/imgcolors"
```

## Usage

Random color generator for the examples.

```go
import "image/color"
import "math/rand"

func randomColors(n uint) []color.Color {
    colors := make([]color.Color, n)
    for i := range colors {
        colors[i] = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
    }
    return colors
}
```

### Horizontal

```go
img := imgcolors.Horizontal(randomColors(17), 900, 70)
```

Example output:

![img](docs/images/horizontal.png)

### Vertical

```go
img := imgcolors.Vertical(randomColors(10), 45, 45*10)
```

Example output:

![img](docs/images/vertical.png)

### Square

```go
import "image/color/palette"

img := imgcolors.Square(palette.Plan9, 900)
```

Example output:

![img](docs/images/square.png)

### Conic

```go
img := imgcolors.Conic(randomColors(17), 800, 800)
```

Example output:

![img](docs/images/conic.png)

### Radial

```go
img := imgcolors.Radial(randomColors(13), 800, 800)
```

Example output:

![img](docs/images/radial.png)

## Playground

- [Try it online](https://play.golang.org/p/pAnlE6zn04K)

