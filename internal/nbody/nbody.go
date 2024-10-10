package nbody

type NBody struct {
	bodies []*Body
}

func NewNBody(bodies []*Body) *NBody {
	return &NBody{bodies: bodies}
}

func (nbody *NBody) GetBodies() []*Body {
	return nbody.bodies
}
