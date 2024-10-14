package nbody

type NBody struct {
	bodies map[int64]*Body
}

func NewNBody(bodies []*Body) *NBody {
	bodiesMap := make(map[int64]*Body, len(bodies))
	for _, body := range bodies {
		bodiesMap[body.Id] = body
	}
	return &NBody{bodies: bodiesMap}
}

func (nbody *NBody) GetBodies() []*Body {
	arr := make([]*Body, len(nbody.bodies))
	for i, body := range nbody.bodies {
		arr[i] = body
	}
	return arr
}

func (nbody *NBody) PopBody(id int64) *Body {
	body := nbody.bodies[id]
	delete(nbody.bodies, id)
	return body
}

func (nbody *NBody) AddBody(body *Body) (*Body, error) {
	if _, ok := nbody.bodies[body.Id]; ok {
		return nil, ErrBodyAlreadyExists
	}
	nbody.bodies[body.Id] = body
	return body, nil
}

func (nbody *NBody) UpdateSelf(uNBody *NBody) *NBody {
	for id, body := range uNBody.bodies {
		nbody.bodies[id].Update(body)
	}
	return nbody
}

func (nbody *NBody) Copy() *NBody {
	newBodies := make(map[int64]*Body, len(nbody.bodies))
	for id, body := range nbody.bodies {
		newBodies[id] = body.Copy()
	}
	return &NBody{bodies: newBodies}
}
