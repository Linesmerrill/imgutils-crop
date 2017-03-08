// Package crop provides image cropping utilities.
package crop

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Anchor specifies the reference point for cropping operations.
type Anchor int

const (
	// Center crops from the center of the image.
	Center Anchor = iota
	// TopLeft crops from the top-left corner.
	TopLeft
	// TopRight crops from the top-right corner.
	TopRight
	// BottomLeft crops from the bottom-left corner.
	BottomLeft
	// BottomRight crops from the bottom-right corner.
	BottomRight
)

// Rectangle crops an image to the specified rectangle.
func Rectangle(src image.Image, rect image.Rectangle) image.Image {
	bounds := src.Bounds()

	// Clamp rectangle to image bounds
	if rect.Min.X < bounds.Min.X {
		rect.Min.X = bounds.Min.X
	}
	if rect.Min.Y < bounds.Min.Y {
		rect.Min.Y = bounds.Min.Y
	}
	if rect.Max.X > bounds.Max.X {
		rect.Max.X = bounds.Max.X
	}
	if rect.Max.Y > bounds.Max.Y {
		rect.Max.Y = bounds.Max.Y
	}

	dst := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))
	draw.Draw(dst, dst.Bounds(), src, rect.Min, draw.Src)
	return dst
}

// ToSize crops an image to the specified size using the given anchor point.
func ToSize(src image.Image, width, height int, anchor Anchor) image.Image {
	bounds := src.Bounds()
	srcW := bounds.Dx()
	srcH := bounds.Dy()

	if width > srcW {
		width = srcW
	}
	if height > srcH {
		height = srcH
	}

	var x, y int
	switch anchor {
	case Center:
		x = (srcW - width) / 2
		y = (srcH - height) / 2
	case TopLeft:
		x, y = 0, 0
	case TopRight:
		x = srcW - width
		y = 0
	case BottomLeft:
		x = 0
		y = srcH - height
	case BottomRight:
		x = srcW - width
		y = srcH - height
	}

	rect := image.Rect(x, y, x+width, y+height)
	return Rectangle(src, rect)
}

// Square crops an image to a square using the given anchor point.
// The size is determined by the smaller dimension.
func Square(src image.Image, anchor Anchor) image.Image {
	bounds := src.Bounds()
	size := bounds.Dx()
	if bounds.Dy() < size {
		size = bounds.Dy()
	}
	return ToSize(src, size, size, anchor)
}

// CenterSquare crops an image to a centered square.
func CenterSquare(src image.Image) image.Image {
	return Square(src, Center)
}

// Margins crops an image by removing the specified margins.
func Margins(src image.Image, top, right, bottom, left int) image.Image {
	bounds := src.Bounds()
	rect := image.Rect(
		bounds.Min.X+left,
		bounds.Min.Y+top,
		bounds.Max.X-right,
		bounds.Max.Y-bottom,
	)
	return Rectangle(src, rect)
}

// CropFromFile reads an image file and crops it.
func CropFromFile(path string, rect image.Rectangle) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return Rectangle(src, rect), nil
}

// SaveJPEG saves the cropped image as JPEG.
func SaveJPEG(img image.Image, w io.Writer, quality int) error {
	if quality <= 0 || quality > 100 {
		quality = 85
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}

// SavePNG saves the cropped image as PNG.
func SavePNG(img image.Image, w io.Writer) error {
	return png.Encode(w, img)
}
