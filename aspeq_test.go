package aspeq

import "testing"

func TestSuper16(t *testing.T) {
	name := FromWidthHeight(5, 3).name
	if name != "super16" {
		t.Error("5:3 should be super16")
	}
}

func TestSquare(t *testing.T) {
	name := FromWidthHeight(100, 102).name
	if name != "square" {
		t.Error("1:1.02 should be square")
	}
}

func TestInstax(t *testing.T) {
	name := FromWidthHeight(46, 62).name
	if name != "instax" {
		t.Errorf("1:1.02 should be instax, was %s", name)
	}
}

func TestExtremelyWide(t *testing.T) {
	name := FromWidthHeight(300, 1).name
	if name != "widelux" {
		t.Errorf("300:1 should be closest to widelux, was %s", name)
	}
}

func TestImage(t *testing.T) {
	ar, _ := FromImage("1.66.jpeg")
	if ar.name != "super16" {
		t.Errorf("1.66.jpeg sould be super16, was %s", ar.name)
	}
}

func TestNonexistentImage(t *testing.T) {
	_, err := FromImage("1.67.jpeg")
	if err == nil {
		t.Errorf("Getting the aspect ratio for 1.67.jpeg should fail")
	}
}
