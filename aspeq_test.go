package aspeq_test

import (
	"blekksprut.net/aspeq"
	"fmt"
	"testing"
)

func TestZeroXy(t *testing.T) {
	xy := aspeq.Match(0, 0).Xy()
	if xy != "12:1" {
		t.Errorf("should return an invalid ratio")
	}
}

func TestXy(t *testing.T) {
	xy := aspeq.Match(5, 3).Xy()
	if xy != "5:3" {
		t.Errorf("5:3 should be 5:3, was %s", xy)
	}
}

func TestOrientation(t *testing.T) {
	or := aspeq.Match(5, 3).Orientation
	if or != aspeq.Landscape {
		t.Error("5:3 should be landscape")
	}
}

func TestSuper16(t *testing.T) {
	name := aspeq.Match(5, 3).Name
	if name != "super16" {
		t.Errorf("5:3 should be super16, was %s", name)
	}
}

func TestSquare(t *testing.T) {
	name := aspeq.Match(100, 102).Name
	if name != "square" {
		t.Errorf("1:1.02 should be square, was %s", name)
	}
}

func TestInstax(t *testing.T) {
	name := aspeq.Match(46, 62).Name
	if name != "instax" {
		t.Errorf("1:1.02 should be instax, was %s", name)
	}
}

func TestExtremelyWide(t *testing.T) {
	name := aspeq.Match(300, 1).Name
	if name != "circle-vision" {
		t.Errorf("300:1 should be closest to circle-vision, was %s", name)
	}
}

func TestImage(t *testing.T) {
	ar, err := aspeq.FromPath("1.66.jpeg")
	if err != nil {
		t.Errorf("something went wrong: %s", err)
		return
	}
	if ar.Name != "super16" {
		t.Errorf("1.66.jpeg should be super16, was %s", ar.Name)
	}
}

func TestBrokenImage(t *testing.T) {
	_, err := aspeq.FromPath("README.md")
	if err == nil {
		t.Errorf("getting the aspect ratio for README.md should fail")
	}
}

func TestNonexistentImage(t *testing.T) {
	_, err := aspeq.FromPath("1.67.jpeg")
	if err == nil {
		t.Errorf("getting the aspect ratio for 1.67.jpeg should fail")
	}
}

func TestNonexistentImageCrop(t *testing.T) {
	_, err := aspeq.CropPath("1.67.jpeg", aspeq.Cinerama)
	if err == nil {
		t.Errorf("cropping an image that doesn't exist should fail")
	}
}

func TestBrokenCrop(t *testing.T) {
	_, err := aspeq.CropPath("README.md", aspeq.Cinerama)
	if err == nil {
		t.Errorf("cropping README.md should fail")
	}
}

func TestRegister(t *testing.T) {
	aspeq.Register("test", 21, 9)
	ar := aspeq.Match(21, 9)
	if ar.Name != "test" {
		t.Errorf("custom aspect ratio not found, got %s", ar.Name)
	}
}

func ExampleMatch() {
	ar := aspeq.Match(1920, 1080)
	fmt.Println(ar.Name)
	// Output: sixteen-nine
}

func ExampleFromPath() {
	ar, err := aspeq.FromPath("1.66.jpeg")
	if err != nil {
		panic(err)
	}
	fmt.Println(ar.Name)
	// Output: super16
}

func ExampleCropPath() {
	img, err := aspeq.CropPath("1.66.jpeg", aspeq.Square)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	fmt.Printf("%dx%d\n", bounds.Dx(), bounds.Dy())
	// Output: 24x24
}

func ExampleCropPath_cinerama() {
	img, err := aspeq.CropPath("1.66.jpeg", aspeq.Cinerama)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	fmt.Printf("%dx%d\n", bounds.Dx(), bounds.Dy())
	// Output: 40x15
}

func ExampleCropPath_classic() {
	img, err := aspeq.CropPath("1.66.jpeg", aspeq.Classic)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	fmt.Printf("%dx%d\n", bounds.Dx(), bounds.Dy())
	// Output: 16x24
}
