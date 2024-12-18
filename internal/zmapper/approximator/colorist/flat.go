package colorist

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"image/color"
)

type FlatColorist struct {
	lights []object.Light
	view   vector.Vector3d
}

type FlatColorModel struct {
	C color.RGBA
}

func NewFlatColorist() *FlatColorist {
	return &FlatColorist{lights: make([]object.Light, 0), view: *vector.NewVector3d(0, 0, 0)}
}

func (c *FlatColorist) VisitPolygonObject(po *object.PolygonObject) {
	for _, p := range po.GetPolygons() {
		normal := p.GetNormal()
		pcolor := p.GetColor()
		color := FlatColorModel{
			C: color.RGBA{0, 0, 0, 255},
		}
		for _, l := range c.lights {
			v1, _, _ := p.GetVertices()
			lcolor := l.CalculateLightContribution(*v1, c.view, *normal, pcolor)
			color.C = mathutils.AddRGBA(color.C, lcolor)
		}
		p.SetColorModel(&color)
	}
}

func (c *FlatColorist) VisitPointLight(light *object.PointLight) {
	c.lights = append(c.lights, light)
}

func (c *FlatColorist) VisitPointLightShadow(light *object.PointLightShadow) {
	c.lights = append(c.lights, light)
}

func (c *FlatColorist) VisitCamera(cam *object.Camera) {
	c.view = cam.GetCenter()
}

func (c *FlatColorist) VisitObjectPool(pool *object.ObjectPool) {
	for _, obj := range pool.GetObjects() {
		obj.Accept(c)
	}
}
