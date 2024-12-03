package plane

import (
	"NBodySim/internal/mathutils/vector"
	"math"
	"testing"
)

func TestNewNormalisation(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	plane.normal.Normalize()
	if plane.normal.Square() != 1 {
		t.Error("Expected normalized normal to have length 1, got", plane.normal.Square())
	}
}

func TestDistance(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	point := *vector.NewVector3d(5, 6, 7)
	distance := plane.Distance(&point)
	if math.Abs(distance-11.2249) > 1e-4 {
		t.Error("Expected distance to be 11.2249, got", distance)
	}
}

func TestSignedDistance(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	point := *vector.NewVector3d(5, 6, 7)
	signedDistance := plane.SignedDistance(&point)
	if signedDistance-11.2249 > 1e-4 {
		t.Error("Expected signed distance to be 11.2249, got", signedDistance)
	}
}

func TestSignedDistanceNegative(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	point := *vector.NewVector3d(-5, -6, -7)
	signedDistance := plane.SignedDistance(&point)
	if !(signedDistance+9.0868 < 1e-4 && signedDistance+9.0868 > -1e-4) {
		t.Error("Expected signed distance to be -9.0868, got", signedDistance)
	}
}

func TestIsPointInFront(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	point := *vector.NewVector3d(5, 6, 7)
	if !plane.IsPointInFront(&point) {
		t.Error("Expected point to be in front of the plane")
	}
}

func TestIsPointBehind(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	point := *vector.NewVector3d(-5, -6, -7)
	if !plane.IsPointBehind(&point) {
		t.Error("Expected point to be behind the plane")
	}
}

func TestIntersection(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	p1 := *vector.NewVector3d(-10, -10, -10)
	p2 := *vector.NewVector3d(10, 10, 10)
	intersection, ok := plane.Intersection(&p1, &p2)
	if !ok {
		t.Error("Expected intersection point")
		return
	}
	if math.Abs(intersection.X+0.67) > 1e-2 || math.Abs(intersection.Y+0.67) > 1e-2 || math.Abs(intersection.Z+0.67) > 1e-2 {
		t.Error("Expected intersection point to be (5, 5, 5), got", intersection)
	}
}

func TestIntersectionAbove(t *testing.T) {
	plane := NewPlane(*vector.NewVector3d(1, 2, 3), 4)
	p1 := *vector.NewVector3d(0, 0, 0)
	p2 := *vector.NewVector3d(10, 10, 10)
	_, ok := plane.Intersection(&p1, &p2)
	if ok {
		t.Error("Expected no intersection")
	}
}
