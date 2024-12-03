package mathutils

import (
	"math"
	"testing"
)

func TestToRadians(t *testing.T) {
	degrees := 45.0
	radians := ToRadians(degrees)
	if radians != math.Pi/4 {
		t.Error("Expected radians to be Pi/4, got", radians)
	}
}

func TestToDegrees(t *testing.T) {
	radians := math.Pi / 4
	degrees := ToDegrees(radians)
	if degrees != 45.0 {
		t.Error("Expected degrees to be 45.0, got", degrees)
	}
}

func TestToRadians0(t *testing.T) {
	degrees := 0.0
	radians := ToRadians(degrees)
	if radians != 0 {
		t.Error("Expected radians to be 0, got", radians)
	}
}

func TestToDegrees0(t *testing.T) {
	radians := 0.0
	degrees := ToDegrees(radians)
	if degrees != 0 {
		t.Error("Expected degrees to be 0, got", degrees)
	}
}
