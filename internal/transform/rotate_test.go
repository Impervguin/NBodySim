package transform

import (
	"NBodySim/internal/mathutils/vector"
	"math"
	"testing"
)

func TestRotateActionX(t *testing.T) {
	rotate := NewRotateAction(vector.NewVector3d(math.Pi/2, 0, 0))
	v := vector.NewVector3d(0, 0, 1)
	rotate.ApplyToVector(v)
	if math.Abs(v.X) > 1e-6 || v.Y != -1 || math.Abs(v.Z) > 1e-6 {
		t.Error("Expected (0, -1, 0), got", v)
	}
}

func TestRotateActionY(t *testing.T) {
	rotate := NewRotateAction(vector.NewVector3d(0, math.Pi/2, 0))
	v := vector.NewVector3d(1, 0, 0)
	rotate.ApplyToVector(v)
	if math.Abs(v.X) > 1e-6 || math.Abs(v.Y) > 1e-6 || v.Z != -1 {
		t.Error("Expected (0, 0, -1), got", v)
	}
}

func TestRotateActionZ(t *testing.T) {
	rotate := NewRotateAction(vector.NewVector3d(0, 0, math.Pi/2))
	v := vector.NewVector3d(1, 0, 0)
	rotate.ApplyToVector(v)
	if math.Abs(v.X) > 1e-6 || v.Y != 1 || math.Abs(v.Z) > 1e-6 {
		t.Error("Expected (1, 0, 0), got", v)
	}
}

func TestRotateActionCenter(t *testing.T) {
	center := vector.NewVector3d(1, 1, 1)
	rotate := NewRotateActionCenter(center, vector.NewVector3d(math.Pi/2, math.Pi/2, 0))
	v := vector.NewVector3d(1, 2, 2)
	rotate.ApplyToVector(v)
	if v.X != 2 || math.Abs(v.Y) > 1e-6 || v.Z != 1 {
		t.Error("Expected (0, -1, 0), got", v)
	}
}

func TestRotateAxisAction(t *testing.T) {
	axis := vector.NewVector3d(0, 1, 1)
	axis.Normalize()
	rotate := NewAxisRotateAction(axis, math.Pi)
	v := vector.NewVector3d(0, 1, 0)
	rotate.ApplyToVector(v)
	if math.Abs(v.X) > 1e-6 || math.Abs(v.Y) > 1e-6 || math.Abs(v.Z-1) > 1e-6 {
		t.Error("Expected (0, 0, 1), got", v)
	}
}
