package gobbox

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

// DrawHorizontalLine draw an horizontal line on img
// from x1 to x2 at y with color as the color
func DrawHorizontalLine(x1, y, x2 int, img draw.Image, color color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, color)
	}
}

// DrawVerticalLine draw a vertical line on img
// from y1 to y2 at x with color as the color
func DrawVerticalLine(x, y1, y2 int, img draw.Image, color color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, color)
	}
}

// DrawRect draw a rectangle on img with coordinates x1, y1, x2, y2,
// with color as the color
func DrawRect(x1, y1, x2, y2 int, img draw.Image, color color.Color) {
	DrawHorizontalLine(x1, y1, x2, img, color)
	DrawHorizontalLine(x1, y2, x2, img, color)
	DrawVerticalLine(x1, y1, y2, img, color)
	DrawVerticalLine(x2, y1, y2, img, color)
}

// WriteLabel write the label at the top left of the img bouding box
// The text color is specified with the color argument
func WriteLabel(img draw.Image, label string, left int, top int, color color.Color) {
	// Load TTF font
	f, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}

	// Create text drawer
	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(color),
		Face: truetype.NewFace(f, &truetype.Options{
			Size: 12,
			DPI:  80,
		}),
		Dot: fixed.P(left+2, top+13),
	}

	// Apply text
	d.DrawString(label)
}

// DrawBoundingBox draw a bouding box on img at x1, x2, y1, y2,
// and if label is specify, it also draw the label on top
// left of the bounding box.
// Color that should be used for the bounding box is specified
// with bboxColor and the color that should be used for the
// text is specified with labelColor
func DrawBoundingBox(img draw.Image, label string, x1, x2, y1, y2 int, bboxColor, labelColor color.Color) {
	// Draw the rectangle
	DrawRect(x1, y1, x2, y2, img, bboxColor)

	// Write the label
	WriteLabel(img, label, x1, y1, labelColor)
}
