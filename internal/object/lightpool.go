package object

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
