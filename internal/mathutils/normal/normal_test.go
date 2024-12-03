package normal

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"testing"
)

func TestNewNormal(t *testing.T) {
	normal := NewNormal(*vector.NewVector3d(1, 2, 3), *vector.NewVector3d(4, 5, 6))
	if normal.Start.X != 1 || normal.Start.Y != 2 || normal.Start.Z != 3 {
		t.Error("Expected (1, 2, 3), got", normal.Start)
	}
	if normal.End.X != 4 || normal.End.Y != 5 || normal.End.Z != 6 {
		t.Error("Expected (4, 5, 6), got", normal.End)
	}
}

func TestToVector(t *testing.T) {
	normal := NewNormal(*vector.NewVector3d(1, 2, 3), *vector.NewVector3d(4, 5, 6))
	expected := *vector.NewVector3d(3, 3, 3)
	if normal.ToVector().X != expected.X || normal.ToVector().Y != expected.Y || normal.ToVector().Z != expected.Z {
		t.Error("Expected (3, 3, 3), got", normal.ToVector())
	}
}

func TestTransform(t *testing.T) {
	normal := NewNormal(*vector.NewVector3d(1, 2, 3), *vector.NewVector3d(4, 5, 6))
	transform := transform.NewMoveAction(vector.NewVector3d(1, 2, 3))
	normal.Transform(transform)
	if normal.Start.X != 2 || normal.Start.Y != 4 || normal.Start.Z != 6 {
		t.Error("Expected (2, 4, 6), got", normal.Start)
	}
	if normal.End.X != 5 || normal.End.Y != 7 || normal.End.Z != 9 {
		t.Error("Expected (5, 7, 9), got", normal.End)
	}
}

func TestCopy(t *testing.T) {
	normal := NewNormal(*vector.NewVector3d(1, 2, 3), *vector.NewVector3d(4, 5, 6))
	copied := normal.Copy()
	if copied.Start.X != 1 || copied.Start.Y != 2 || copied.Start.Z != 3 {
		t.Error("Expected (1, 2, 3), got", copied.Start)
	}
	if copied.End.X != 4 || copied.End.Y != 5 || copied.End.Z != 6 {
		t.Error("Expected (4, 5, 6), got", copied.End)
	}
	if normal == copied {
		t.Error("Expected copied normal to be different from the original, got same reference")
	}
}
