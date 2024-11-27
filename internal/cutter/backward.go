package cutter

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
)

type BackwardsCutter struct {
	camera  *object.Camera
	forward vector.Vector3d
}

func NewBackwardsCutter(camera *object.Camera) *BackwardsCutter {
	return &BackwardsCutter{
		camera:  camera,
		forward: camera.GetForward(),
	}
}

func (c *BackwardsCutter) VisitCamera(camera *object.Camera) {
	c.camera = camera
	c.forward = camera.GetForward()
}

func (c *BackwardsCutter) VisitPolygonObject(po *object.PolygonObject) {
	polygons := po.GetPolygons()
	po.ResetVertices()
	po.ResetPolygons()

	vertMap := make(map[*vector.Vector3d]struct{}, 2*len(polygons))
	for _, polygon := range polygons {
		if c.SeePolygon(polygon) {
			po.AddPolygon(polygon)
			v1, v2, v3 := polygon.GetVertices()
			vertMap[v1] = struct{}{}
			vertMap[v2] = struct{}{}
			vertMap[v3] = struct{}{}
		}
	}
	for v := range vertMap {
		po.AddVertex(v)
	}
}

func (c *BackwardsCutter) SeePolygon(polygon *object.Polygon) bool {
	normal := polygon.GetNormal().ToVector()
	if polygon.NormalIsOuter() && normal.Dot(&c.forward) >= -1e-6 {
		return false
	}
	if polygon.NormalIsInner() && normal.Dot(&c.forward) <= 1e-6 {
		return false
	}
	return true
}

func (c *BackwardsCutter) VisitObjectPool(pool *object.ObjectPool) {
	for _, obj := range pool.GetObjects() {
		obj.Accept(c)
	}
}

func SeePoint(point *vector.Vector3d) bool {
	return true
}
