package colorist

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"image/color"
)

type PolygonColor struct {
	C1 color.RGBA
	C2 color.RGBA
	C3 color.RGBA
}

type GuroColorModel struct {
	C map[int64]PolygonColor
}

type GuroColorist struct {
	view   vector.Vector3d
	lights []object.Light
}

func NewGuroColorist() *GuroColorist {
	return &GuroColorist{view: *vector.NewVector3d(0, 0, 0), lights: make([]object.Light, 0)}
}

func (c *GuroColorist) VisitPolygonObject(po *object.PolygonObject) {
	for _, p := range po.GetPolygons() {
		c.processPolygon(p)
	}
}

func (c *GuroColorist) processPolygon(p *object.Polygon) {
	normal := p.GetNormal()
	pcolor := p.GetColor()
	color := GuroColorModel{
		C: make(map[int64]PolygonColor, len(c.lights)),
	}
	v1, v2, v3 := p.GetVertices()
	for _, l := range c.lights {
		l1color := l.CalculateLightContribution(*v1, c.view, *normal, pcolor)
		l2color := l.CalculateLightContribution(*v2, c.view, *normal, pcolor)
		l3color := l.CalculateLightContribution(*v3, c.view, *normal, pcolor)
		color.C[l.GetId()] = PolygonColor{mathutils.ToRGBA(l1color), mathutils.ToRGBA(l2color), mathutils.ToRGBA(l3color)}
	}
	p.SetColorModel(&color)
}

func (c *GuroColorist) VisitPointLight(light *object.PointLight) {
	c.lights = append(c.lights, light)
}

func (c *GuroColorist) VisitPointLightShadow(light *object.PointLightShadow) {
	c.lights = append(c.lights, light)
}

func (c *GuroColorist) VisitCamera(cam *object.Camera) {
	c.view = cam.GetCenter()
}

func (c *GuroColorist) VisitObjectPool(pool *object.ObjectPool) {
	for _, obj := range pool.GetObjects() {
		obj.Accept(c)
	}
}
