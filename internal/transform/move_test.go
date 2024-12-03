package transform

import (
	"NBodySim/internal/mathutils/vector"
	"testing"
)

func TestMoveAction(t *testing.T) {
	translation := vector.NewVector3d(1, 2, 3)
	move := NewMoveAction(translation)
	homo := vector.NewHomoVector(4, 5, 6, 1)
	move.ApplyToHomoVector(homo)
	if homo.X != 5 || homo.Y != 7 || homo.Z != 9 {
		t.Error("Expected (5, 7, 9), got", homo)
	}
}
