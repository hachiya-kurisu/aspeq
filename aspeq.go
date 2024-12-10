package aspeq

import (
	"github.com/anthonynsimon/bild/imgio"
	"image"
	"math"
)

const Version = "0.0.1"

type AspectRatio struct {
	ratio float64
	name  string
}

var Ratios = [...]AspectRatio{
	{0.5625, "tiktok"},
	{0.75, "instax"},
	{1.0, "square"},
	{1.19, "movietone"},
	{1.333, "four-thirds"},
	{1.375, "academy"},
	{1.50, "35mm"},
	{1.66, "super16"},
	{1.77, "sixteen-nine"},
	{1.85, "flat"},
	{2.35, "cinemascope"},
	{2.59, "cinerama"},
	{3.0, "widelux"},
}

func FromWidthHeight(w int, h int) string {
	ratio := float64(w) / float64(h)
	current := Ratios[0]
	for _, candidate := range Ratios {
		if math.Abs(ratio-candidate.ratio) > math.Abs(ratio-current.ratio) {
			return current.name
		}
		current = candidate
	}
	return current.name
}

func FromRectangle(r image.Rectangle) string {
	return FromWidthHeight(r.Dx(), r.Dy())
}

func FromImage(path string) (string, error) {
	img, err := imgio.Open(path)
	if err != nil {
		return "", err
	}
	return FromRectangle(img.Bounds()), nil
}
