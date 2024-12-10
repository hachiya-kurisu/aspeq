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

