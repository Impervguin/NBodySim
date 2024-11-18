package object

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"container/list"
	"fmt"
)

type PolygonObject struct {
	VisibleObject
	ObjectWithId
	vertices *list.List
	polygons *list.List
	center   vector.Vector3d
}

func NewPolygonObject(vertices []*vector.Vector3d, polygons []*Polygon, center vector.Vector3d) *PolygonObject {
	po := &PolygonObject{
		vertices: list.New(),
		polygons: list.New(),
		center:   center,
	}
	po.id = getNextId()
	for _, vertex := range vertices {
		po.AddVertex(vertex)
	}
	for _, polygon := range polygons {
		po.AddPolygon(polygon)
	}
	return po
}

func (po *PolygonObject) ResetVertices() {
	po.vertices = list.New()
}

func (po *PolygonObject) ResetPolygons() {
	po.polygons = list.New()
}

func (po *PolygonObject) AddVertex(vertex *vector.Vector3d) *vector.Vector3d {
	inst, _ := (po.vertices.PushBack(vertex).Value).(*vector.Vector3d)
	return inst
}

func (po *PolygonObject) AddPolygon(polygon *Polygon) *Polygon {
	inst, _ := (po.polygons.PushBack(polygon).Value).(*Polygon)
	return inst
}

func (po *PolygonObject) GetCenter() vector.Vector3d {
	return po.center
}

func (po *PolygonObject) GetVertices() []*vector.Vector3d {
	res := make([]*vector.Vector3d, po.vertices.Len())
	i := 0
	for el := po.vertices.Front(); el != nil; el = el.Next() {
		res[i] = el.Value.(*vector.Vector3d)
		i++
	}
	return res
}

func (po *PolygonObject) GetPolygons() []*Polygon {
	res := make([]*Polygon, po.polygons.Len())
	i := 0
	for el := po.polygons.Front(); el != nil; el = el.Next() {
		res[i] = el.Value.(*Polygon)
		i++
	}
	return res
}

func (po *PolygonObject) Clone() Object {
	vertices := make([]*vector.Vector3d, po.vertices.Len())

	overtices := po.GetVertices()
	vertMap := make(map[*vector.Vector3d]int, po.vertices.Len())

	for i, vertex := range overtices {
		vertices[i] = vertex.Copy()
		vertMap[vertex] = i
	}

	polygons := make([]*Polygon, po.polygons.Len())
	i := 0
	for el := po.polygons.Front(); el != nil; el = el.Next() {
		polygon := el.Value.(*Polygon)
		polygons[i] = polygon.Clone()
		polygons[i].v1 = vertices[vertMap[polygon.v1]]
		polygons[i].v2 = vertices[vertMap[polygon.v2]]
		polygons[i].v3 = vertices[vertMap[polygon.v3]]
		polygons[i].color = polygon.color
		i++
	}
	return NewPolygonObject(vertices, polygons, po.center)
}

func (po *PolygonObject) Transform(action transform.TransformAction) {
	for el := po.vertices.Front(); el != nil; el = el.Next() {
		vertex := el.Value.(*vector.Vector3d)
		action.ApplyToVector(vertex)
	}
	for el := po.polygons.Front(); el != nil; el = el.Next() {
		polygon := el.Value.(*Polygon)
		polygon.TransformNormal(action)
	}
	action.ApplyToVector(&po.center)
}

func (po *PolygonObject) PrintPoints() {
	for el := po.vertices.Front(); el != nil; el = el.Next() {
		vertex := el.Value.(*vector.Vector3d)
		fmt.Printf("Point: %+v\n", vertex)
	}
}

func (po *PolygonObject) PrintPolygons() {
	for el := po.polygons.Front(); el != nil; el = el.Next() {
		polygon := el.Value.(*Polygon)
		fmt.Printf("Polygon: %+v %+v %+v\n", polygon.v1, polygon.v2, polygon.v3)
	}
}

func (po *PolygonObject) Accept(visitor ObjectVisitor) {
	visitor.VisitPolygonObject(po)
}
