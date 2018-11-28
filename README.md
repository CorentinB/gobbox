[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![Build Status](https://travis-ci.org/CorentinB/gobbox.svg?branch=master)](https://travis-ci.org/CorentinB/gobbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/CorentinB/gobbox)](https://goreportcard.com/report/github.com/CorentinB/gobbox)
[![GoDoc](https://godoc.org/github.com/CorentinB/gobbox?status.svg)](https://godoc.org/github.com/CorentinB/gobbox)

# GoBBox

:white_square_button: Pure Go bounding boxes generation with labeling

# Install

`go get -u github.com/CorentinB/gobbox`

# Examples

Simple example of applying a bounding box and a label to an image.

```go
// Colors declarations
red := color.RGBA{255, 0, 0, 255}
white := color.RGBA{255, 255, 255, 255}

// Create a 300x300 image initialized to all-blue
img := image.NewRGBA(image.Rect(0, 0, 640, 480))
blue := color.RGBA{0, 0, 255, 255}
draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

// Write a red bounding box with the "gobbox" label with white font
// at coordinates 100, 250, 238, 68
gobbox.DrawBoundingBox(img, "gobbox", 100, 250, 238, 68, red, white)
```