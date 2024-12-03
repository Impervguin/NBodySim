package mathutils

import (
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {
	f := 4.7
	expected := 5
	result := ToInt(f)
	if result != expected {
		t.Error("Expected 5, got", result)
	}
}

func TestToInt2(t *testing.T) {
	f := -4.2
	expected := -4
	result := ToInt(f)
	if result != expected {
		t.Error("Expected -4, got", result)
	}
}

func TestIAbs(t *testing.T) {
	a := 5
	expected := 5
	result := IAbs(a)
	if result != expected {
		t.Error("Expected 5, got", result)
	}
	a = -3
	expected = 3
	result = IAbs(a)
	if result != expected {
		t.Error("Expected 3, got", result)
	}
}

func TestLinearXIntInterpolation(t *testing.T) {
	x1 := 0
	y1 := 0
	x2 := 5
	y2 := 10
	expected := [][]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}, {5, 10}}
	result := LinearXIntInterpolation(x1, y1, x2, y2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}

func TestLinearYIntInterpolation(t *testing.T) {
	x1 := 0
	y1 := 0
	x2 := 5
	y2 := 10
	expected := [][]int{{0, 0}, {1, 1}, {1, 2}, {2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}, {4, 8}, {5, 9}, {5, 10}}
	result := LinearYIntInterpolation(x1, y1, x2, y2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}

func TestLinearXInterpolation(t *testing.T) {
	x1 := 0
	y1 := 0.0
	x2 := 5
	y2 := 10.0
	expected := []float64{0.0, 2.0, 4.0, 6.0, 8.0, 10.0}
	_, result := LinearXInterpolation(x1, y1, x2, y2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}

func TestLinearXInterpolationSwap(t *testing.T) {
	x1 := 5
	y1 := 0.0
	x2 := 0
	y2 := 10.0
	expected := []float64{10.0, 8, 6, 4, 2, 0}
	_, result := LinearXInterpolation(x1, y1, x2, y2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}
