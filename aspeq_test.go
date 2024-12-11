package aspeq

import "testing"

func TestXy(t *testing.T) {
	xy := FromWidthHeight(5, 3).Xy()
	if xy != "5:3" {
		t.Errorf("5:3 should be 5:3, was %s", xy)
	}
}

func TestOrientation(t *testing.T) {
	or := FromWidthHeight(5, 3).Orientation
	if or != Landscape {
		t.Error("5:3 should be landscape")
	}
}

func TestSuper16(t *testing.T) {
	name := FromWidthHeight(5, 3).Name
	if name != "super16" {
		t.Errorf("5:3 should be super16, was %s", name)
	}
}

func TestSquare(t *testing.T) {
	name := FromWidthHeight(100, 102).Name
	if name != "square" {
		t.Errorf("1:1.02 should be square, was %s", name)
	}
}

func TestInstax(t *testing.T) {
	name := FromWidthHeight(46, 62).Name
	if name != "instax" {
		t.Errorf("1:1.02 should be instax, was %s", name)
	}
}

func TestExtremelyWide(t *testing.T) {
	name := FromWidthHeight(300, 1).Name
	if name != "widelux" {
		t.Errorf("300:1 should be closest to widelux, was %s", name)
	}
}

func TestImage(t *testing.T) {
	ar, _ := FromImage("1.66.jpeg")
	if ar.Name != "super16" {
		t.Errorf("1.66.jpeg sould be super16, was %s", ar.Name)
	}
}

func TestNonexistentImage(t *testing.T) {
	_, err := FromImage("1.67.jpeg")
	if err == nil {
		t.Errorf("Getting the aspect ratio for 1.67.jpeg should fail")
	}
}
