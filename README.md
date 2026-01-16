# imgutils-crop

[![Go Reference](https://pkg.go.dev/badge/github.com/imgutils-org/imgutils-crop.svg)](https://pkg.go.dev/github.com/imgutils-org/imgutils-crop)
[![Go Report Card](https://goreportcard.com/badge/github.com/imgutils-org/imgutils-crop)](https://goreportcard.com/report/github.com/imgutils-org/imgutils-crop)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A flexible Go library for cropping images with multiple anchor points and methods. Part of the [imgutils](https://github.com/imgutils-org) collection.

## Features

- Rectangle-based cropping with precise coordinates
- Anchor-based cropping (center, corners)
- Square cropping for profile pictures
- Margin-based cropping
- Automatic bounds clamping

## Installation

```bash
go get github.com/imgutils-org/imgutils-crop
```

## Quick Start

```go
package main

import (
    "image"
    "log"
    "os"

    "github.com/imgutils-org/imgutils-crop"
)

func main() {
    // Open image
    file, _ := os.Open("photo.jpg")
    defer file.Close()
    src, _, _ := image.Decode(file)

    // Crop a centered square (perfect for profile pictures)
    cropped := crop.CenterSquare(src)

    // Save result
    out, _ := os.Create("profile.jpg")
    defer out.Close()
    crop.SaveJPEG(cropped, out, 90)
}
```

## Usage Examples

### Rectangle Crop

```go
// Crop a specific region (x1, y1, x2, y2)
rect := image.Rect(100, 100, 500, 400)
cropped := crop.Rectangle(src, rect)
```

### Anchor-Based Cropping

```go
// Crop 400x300 from the center
cropped := crop.ToSize(src, 400, 300, crop.Center)

// Crop from top-left corner
cropped := crop.ToSize(src, 400, 300, crop.TopLeft)

// Crop from bottom-right corner
cropped := crop.ToSize(src, 400, 300, crop.BottomRight)
```

### Square Cropping

```go
// Centered square crop (great for avatars)
cropped := crop.CenterSquare(src)

// Square from specific anchor
cropped := crop.Square(src, crop.TopLeft)
cropped := crop.Square(src, crop.BottomRight)
```

### Margin-Based Cropping

```go
// Remove 50px from all sides
cropped := crop.Margins(src, 50, 50, 50, 50)

// Remove different margins (top, right, bottom, left)
cropped := crop.Margins(src, 100, 50, 100, 50)
```

### Crop from File

```go
rect := image.Rect(0, 0, 800, 600)
cropped, err := crop.CropFromFile("landscape.jpg", rect)
if err != nil {
    log.Fatal(err)
}
```

## API Reference

### Types

#### Anchor

```go
type Anchor int

const (
    Center      Anchor = iota // Crop from center
    TopLeft                   // Crop from top-left
    TopRight                  // Crop from top-right
    BottomLeft                // Crop from bottom-left
    BottomRight               // Crop from bottom-right
)
```

### Functions

| Function | Description |
|----------|-------------|
| `Rectangle(src, rect)` | Crop to exact rectangle coordinates |
| `ToSize(src, w, h, anchor)` | Crop to size from anchor point |
| `Square(src, anchor)` | Crop to square from anchor |
| `CenterSquare(src)` | Crop to centered square |
| `Margins(src, t, r, b, l)` | Crop by removing margins |
| `CropFromFile(path, rect)` | Load and crop from file |
| `SaveJPEG(img, w, quality)` | Save as JPEG |
| `SavePNG(img, w)` | Save as PNG |

## Common Use Cases

### Profile Pictures

```go
// Create a centered square crop for profile pictures
profile := crop.CenterSquare(src)
```

### Banner Images

```go
// Crop to 16:9 aspect ratio from center
width := src.Bounds().Dx()
height := width * 9 / 16
banner := crop.ToSize(src, width, height, crop.Center)
```

### Remove Borders

```go
// Remove 10px border from all sides
cleaned := crop.Margins(src, 10, 10, 10, 10)
```

## Requirements

- Go 1.16 or later

## Related Packages

- [imgutils-resize](https://github.com/imgutils-org/imgutils-resize) - Image resizing
- [imgutils-thumbnail](https://github.com/imgutils-org/imgutils-thumbnail) - Thumbnail generation
- [imgutils-sdk](https://github.com/imgutils-org/imgutils-sdk) - Unified SDK

## License

MIT License - see [LICENSE](LICENSE) for details.
