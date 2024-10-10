package nbody

type calculateBody func(body *Body, dt float64) (*Body, error)

type NBodyEngine interface {
	Calculate(nbody *NBody, fn calculateBody, dt float64) (*NBody, error)
}

type IterativeNbodyEngine struct{}

func NewIterativeNbodyEngine() *IterativeNbodyEngine {
	return &IterativeNbodyEngine{}
}

func (e *IterativeNbodyEngine) Calculate(nbody *NBody, fn calculateBody, dt float64) (*NBody, error) {
	uNBody := NewNBody(make([]*Body, 0, len(nbody.bodies)))
	for _, body := range nbody.bodies {
		ubody, err := fn(body, dt)
		if err != nil {
			return nil, err
		}
		uNBody.bodies = append(uNBody.bodies, ubody)
	}
	return uNBody, nil
}
