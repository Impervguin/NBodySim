package object

import (
	"NBodySim/internal/transform"
	"NBodySim/internal/vectormath"
	"image/color"
)

type Polygon struct {
	v1, v2, v3 *vectormath.Vector3d
	color      color.Color
}

type PolygonObject struct {
	VisibleObject
	ObjectWithId
	vertices []vectormath.Vector3d
	polygons []Polygon
	center   vectormath.Vector3d
}

func NewPolygonObject(vertices []vectormath.Vector3d, polygons []Polygon, center vectormath.Vector3d) *PolygonObject {
	po := &PolygonObject{
		vertices: vertices,
		polygons: polygons,
		center:   center,
	}
	po.id = getNextId()
	return po
}

func (po *PolygonObject) GetCenter() vectormath.Vector3d {
	return po.center
}

func (po *PolygonObject) Clone() Object {
	vertices := make([]vectormath.Vector3d, len(po.vertices))
	copy(vertices, po.vertices)
	polygons := make([]Polygon, len(po.polygons))
	for i, polygon := range po.polygons {
		for j := range po.vertices {
			if polygon.v1 == &vertices[j] {
				polygons[i].v1 = &vertices[j]
			}
			if polygon.v2 == &vertices[j] {
				polygons[i].v2 = &vertices[j]
			}
			if polygon.v3 == &vertices[j] {
				polygons[i].v3 = &vertices[j]
			}
		}
	}
	return NewPolygonObject(vertices, polygons, po.center)
}

func (po *PolygonObject) Transform(action *transform.TransformAction) {
	for i := range po.vertices {
		(*action).ApplyToVector(&po.vertices[i])
	}
}

func (po *PolygonObject) Accept(visitor ObjectVisitor) {
	visitor.VisitPolygonObject(po)
}
