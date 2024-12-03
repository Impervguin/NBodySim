package plane

import (
	"NBodySim/internal/mathutils/vector"
	"math"
)

type Plane struct {
	normal vector.Vector3d
	d      float64
}

func NewPlane(normal vector.Vector3d, d float64) *Plane {
	coeff := vector.Length(&normal)
	normal.Normalize()
	return &Plane{normal: normal, d: d / coeff}
}

func (p *Plane) Distance(point *vector.Vector3d) float64 {
	return math.Abs(p.normal.Dot(point) + p.d)
}

func (p *Plane) SignedDistance(point *vector.Vector3d) float64 {
	return p.normal.Dot(point) + p.d
}

func (p *Plane) IsPointInFront(point *vector.Vector3d) bool {
	return p.SignedDistance(point) > 0
}

func (p *Plane) IsPointBehind(point *vector.Vector3d) bool {
	return p.SignedDistance(point) < 0
}

func (p *Plane) Intersection(p1 *vector.Vector3d, p2 *vector.Vector3d) (*vector.Vector3d, bool) {
	s1 := p.SignedDistance(p1)
	s2 := p.SignedDistance(p2)
	if s1 > 0 && s2 > 0 || s1 < 0 && s2 < 0 {
		return nil, false
	}
	t := -s1 / (s2 - s1)
	if t < 0 || t > 1 {
		return nil, false
	}

	return vector.AddVectors(p1.Copy(), vector.MultiplyVectorScalar(vector.SubtractVectors(p2, p1), t)), true

}
