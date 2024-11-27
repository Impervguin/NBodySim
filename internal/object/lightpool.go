package object

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/normal"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"image/color"
)

type LightPoolVisitor interface {
	VisitLightPool(lp *LightPool)
}

type LightPool struct {
	lights map[int64]Light
}

func NewLightPool() *LightPool {
	return &LightPool{lights: make(map[int64]Light)}
}

func (lp *LightPool) GetLight(id int64) (Light, bool) {
	light, found := lp.lights[id]
	return light, found
}

func (lp *LightPool) PutLight(light Light) {
	lp.lights[light.GetId()] = light
}

func (lp *LightPool) RemoveLight(id int64) {
	delete(lp.lights, id)
}

func (lp *LightPool) Accept(visitor LightVisitor) {
	for _, light := range lp.lights {
		light.Accept(visitor)
	}
}

func (lp *LightPool) GetCount() int {
	return len(lp.lights)
}

func (lp *LightPool) Clone() *LightPool {
	newPool := NewLightPool()
	for _, light := range lp.lights {
		newPool.PutLight(light.Clone())
	}
	return newPool
}

func (lp *LightPool) Transform(action transform.TransformAction) {
	for _, light := range lp.lights {
		light.Transform(action)
	}
}

// func (lp *LightPool) Accept(vis LightPoolVisitor) {
// 	vis.VisitLightPool(lp)
// }

func (lp *LightPool) CalculateLight(point, view vector.Vector3d, normal normal.Normal, col color.Color) color.Color {
	contribution := mathutils.ToRGBA64(color.Black)
	for _, light := range lp.lights {
		contribution = mathutils.AddRGBA64(contribution, light.CalculateLightContribution(point, view, normal, col))
	}
	return contribution
}
