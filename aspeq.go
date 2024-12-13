package aspeq

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"image"
	"math"
)

const Version = "0.1.0"

type Orientation int

const (
	Square Orientation = iota + 1
	Portrait
	Landscape
)

type AspectRatio struct {
	Ratio       float64
	Name        string
	X           int64
	Y           int64
	Orientation Orientation
}

var Ratios = [...]AspectRatio{
	{0.5625, "tiktok", 9, 16, Portrait},
	{0.75, "instax", 3, 4, Portrait},
	{1.0, "square", 1, 1, Square},
	{1.19, "movietone", 19, 16, Landscape},
	{1.333, "four-thirds", 4, 3, Landscape},
	{1.375, "academy", 4, 3, Landscape},
	{1.50, "leica", 3, 2, Landscape},
	{1.66, "super16", 5, 3, Landscape},
	{1.77, "sixteen-nine", 16, 9, Landscape},
	{1.85, "flat", 37, 20, Landscape},
	{2.0, "univisium", 2, 1, Landscape},
	{2.35, "cinemascope", 47, 20, Landscape},
	{2.59, "cinerama", 70, 27, Landscape},
	{3.0, "widelux", 3, 1, Landscape},
	{4.0, "polyvision", 4, 1, Landscape},
	{12.0, "circle-vision", 12, 1, Landscape},
}

func (ar AspectRatio) Xy() string {
	return fmt.Sprintf("%d:%d", ar.X, ar.Y)
}

func FromWidthHeight(w int, h int) *AspectRatio {
	ratio := float64(w) / float64(h)
	current := Ratios[0]
	for _, candidate := range Ratios {
		if math.Abs(ratio-candidate.Ratio) > math.Abs(ratio-current.Ratio) {
			return &current
		}
		current = candidate
	}
	return &current
}

func FromRectangle(r image.Rectangle) *AspectRatio {
	return FromWidthHeight(r.Dx(), r.Dy())
}

func FromImage(path string) (*AspectRatio, error) {
	img, err := imgio.Open(path)
	if err != nil {
		return &AspectRatio{}, err
	}
	return FromRectangle(img.Bounds()), nil
}
