package transform

import (
	"NBodySim/internal/mathutils/vector"
	"testing"
)

func TestScaleAction(t *testing.T) {
	scale := NewScaleAction(vector.NewVector3d(2, 3, 4))
	v := vector.NewVector3d(1, 2, 3)
	scale.ApplyToVector(v)
	if v.X != 2 || v.Y != 6 || v.Z != 12 {
		t.Error("Expected (2, 6, 12), got", v)
	}
}

func TestScaleActionCenter(t *testing.T) {
	scale := NewScaleActionCenter(vector.NewVector3d(1, 2, 3), vector.NewVector3d(2, 3, 4))
	v := vector.NewVector3d(1, 2, 3)
	scale.ApplyToVector(v)
	if v.X != 1 || v.Y != 2 || v.Z != 3 {
		t.Error("Expected (1, 2, 3), got", v)
	}
}

func TestScaleActionCenter2(t *testing.T) {
	scale := NewScaleActionCenter(vector.NewVector3d(1, 1, 1), vector.NewVector3d(3, 3, 3))
	v := vector.NewVector3d(1, 2, 3)
	scale.ApplyToVector(v)
	if v.X != 1 || v.Y != 4 || v.Z != 7 {
		t.Error("Expected (1, 2, 3), got", v)
	}
}
