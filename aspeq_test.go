package aspeq

import "testing"

func TestSuper16(t *testing.T) {
	ratio := FromWidthHeight(5, 3)
	if ratio != "super16" {
		t.Error("5:3 should be super16")
	}
}

func TestSquare(t *testing.T) {
	ratio := FromWidthHeight(100, 102)
	if ratio != "square" {
		t.Error("1:1.02 should be square")
	}
}

func TestInstax(t *testing.T) {
	ratio := FromWidthHeight(46, 62)
	if ratio != "instax" {
		t.Errorf("1:1.02 should be instax, was %s", ratio)
	}
}

func TestExtremelyWide(t *testing.T) {
	ratio := FromWidthHeight(300, 1)
	if ratio != "widelux" {
		t.Errorf("300:1 should be closest to widelux, was %s", ratio)
	}
}

func TestImage(t *testing.T) {
	ratio, _ := FromImage("1.66.jpeg")
	if ratio != "super16" {
		t.Errorf("1.66.jpeg sould be super16, was %s", ratio)
	}
}

func TestNonexistentImage(t *testing.T) {
	_, err := FromImage("1.67.jpeg")
	if err == nil {
		t.Errorf("Getting the ratio for 1.67.jpeg should fail")
	}
}
