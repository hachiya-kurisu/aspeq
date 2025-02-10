package aspeq

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"

	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

const Version = "0.3.8"

type Orientation int

const (
	Balanced Orientation = iota + 1
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

var Insta = &AspectRatio{0.5625, "insta", 9, 16, Portrait}
var Classic = &AspectRatio{0.6667, "classic", 2, 3, Portrait}
var Instax = &AspectRatio{0.75, "instax", 3, 4, Portrait}
var Square = &AspectRatio{1.0, "square", 1, 1, Balanced}
var Movietone = &AspectRatio{1.19, "movietone", 19, 16, Landscape}
var FourThirds = &AspectRatio{1.333, "four-thirds", 4, 3, Landscape}
var Academy = &AspectRatio{1.375, "academy", 4, 3, Landscape}
var Leica = &AspectRatio{1.50, "leica", 3, 2, Landscape}
var Super16 = &AspectRatio{1.66, "super16", 5, 3, Landscape}
var SixteenNine = &AspectRatio{1.77, "sixteen-nine", 16, 9, Landscape}
var Flat = &AspectRatio{1.85, "flat", 37, 20, Landscape}
var Univisium = &AspectRatio{2.0, "univisium", 2, 1, Landscape}
var Cinemascope = &AspectRatio{2.35, "cinemascope", 47, 20, Landscape}
var Cinerama = &AspectRatio{2.59, "cinerama", 70, 27, Landscape}
var Widelux = &AspectRatio{3.0, "widelux", 3, 1, Landscape}
var Polyvision = &AspectRatio{4.0, "polyvision", 4, 1, Landscape}
var CircleVision = &AspectRatio{12.0, "circle-vision", 12, 1, Landscape}

var Ratios = [...]*AspectRatio{
	Insta, Classic, Instax, Square, Movietone, FourThirds, Academy,
	Leica, Super16, SixteenNine, Flat, Univisium, Cinemascope, Cinerama,
	Widelux, Polyvision, CircleVision,
}

func (ar *AspectRatio) Xy() string {
	return fmt.Sprintf("%d:%d", ar.X, ar.Y)
}

func Match(w int, h int) *AspectRatio {
	ratio := float64(w) / float64(h)
	current := Ratios[0]
	for _, candidate := range Ratios {
		if math.Abs(ratio-candidate.Ratio) > math.Abs(ratio-current.Ratio) {
			return current
		}
		current = candidate
	}
	return current
}

func FromImage(img image.Image) *AspectRatio {
	bounds := img.Bounds()
	return Match(bounds.Dx(), bounds.Dy())
}

func FromPath(path string) (*AspectRatio, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return FromImage(img), nil
}

func CropImage(img image.Image, ar *AspectRatio) image.Image {
	var w, h, x, y int
	bounds := img.Bounds()
	if float64(bounds.Dx())/float64(bounds.Dy()) > ar.Ratio {
		w = int(float64(bounds.Dy()) * ar.Ratio)
		h = bounds.Dy()
		x = (bounds.Dx() - w) / 2
		y = 0
	} else {
		w = bounds.Dx()
		h = int(float64(bounds.Dx()) / ar.Ratio)
		x = 0
		y = (bounds.Dy() - h) / 2
	}
	cropped := image.NewRGBA(image.Rect(0, 0, w, h))
	src := image.Rect(x, y, x+w, y+h)
	draw.Draw(cropped, cropped.Bounds(), img, src.Min, draw.Over)
	return cropped
}

func CropPath(path string, ar *AspectRatio) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return CropImage(img, ar), nil
}
