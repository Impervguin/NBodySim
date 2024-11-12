package colorist

import (
	"NBodySim/internal/object"
)

type DepthColorist struct {
}

type DepthColorModel struct {
}

func NewDepthColoris() *DepthColorist {
	return &DepthColorist{}
}

func (c *DepthColorist) VisitPolygonObject(po *object.PolygonObject) {
	for _, p := range po.GetPolygons() {
		p.SetColorModel(&DepthColorModel{})
	}
}

func (c *DepthColorist) VisitPointLight(light *object.PointLight) {
}

func (c *DepthColorist) VisitCamera(cam *object.Camera) {
}

func (c *DepthColorist) VisitObjectPool(pool *object.ObjectPool) {
	for _, obj := range pool.GetObjects() {
		obj.Accept(c)
	}
}
