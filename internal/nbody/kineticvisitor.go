package nbody

type KineticVisitor struct {
	kineticEnergy float64
}

func NewKineticVisitor() *KineticVisitor {
	return &KineticVisitor{
		kineticEnergy: 0,
	}
}

func (v *KineticVisitor) VisitBody(body Body) {
	vel := body.GetVelocity()
	v.kineticEnergy += 0.5 * body.GetMass() * vel.Square()
}

func (v *KineticVisitor) VisitNBody(nbody *NBody) {
	for _, b := range nbody.nbody {
		v.VisitBody(b)
	}
}

func (v *KineticVisitor) GetKineticEnergy() float64 {
	return v.kineticEnergy
}
