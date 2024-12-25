package aspeq_test

import (
	"blekksprut.net/aspeq"
	"fmt"
	"testing"
)

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
	ar, err := aspeq.FromImage("1.66.jpeg")
	if err != nil {
		t.Errorf("something went wrong: %s", err)
		return
	}
	if ar.Name != "super16" {
		t.Errorf("1.66.jpeg should be super16, was %s", ar.Name)
	}
}

func TestBrokenImage(t *testing.T) {
	_, err := aspeq.FromImage("README.md")
	if err == nil {
		t.Errorf("getting the aspect ratio for README.md should fail")
	}
}

func TestNonexistentImage(t *testing.T) {
	_, err := aspeq.FromImage("1.67.jpeg")
	if err == nil {
		t.Errorf("getting the aspect ratio for 1.67.jpeg should fail")
	}
}

func ExampleMatch() {
	ar := aspeq.Match(1920, 1080)
	fmt.Println(ar.Name)
	// Output: sixteen-nine
}

func ExampleFromImage() {
	ar, err := aspeq.FromImage("1.66.jpeg")
	if err != nil {
		panic(err)
	}
	fmt.Println(ar.Name)
	// Output: super16
}

