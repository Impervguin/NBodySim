package nbody

type EnergyVisitor struct {
	kineticEnergy   float64
	potentialEnergy float64
}

func NewEnergyVisitor() *EnergyVisitor {
	return &EnergyVisitor{
		kineticEnergy:   0,
		potentialEnergy: 0,
	}
}

func (v *EnergyVisitor) VisitBody(body Body) {
	k := NewKineticVisitor()
	k.VisitBody(body)
	v.kineticEnergy += k.GetKineticEnergy()
}

func (v *EnergyVisitor) VisitNBody(nbody *NBody) {
	p := NewPotentialVisitor()
	k := NewKineticVisitor()
	p.VisitNBody(nbody)
	k.VisitNBody(nbody)
	v.potentialEnergy += p.GetPotentialEnergy()
	v.kineticEnergy += k.GetKineticEnergy()
}

func (v *EnergyVisitor) GetTotalEnergy() float64 {
	return v.kineticEnergy + v.potentialEnergy
}
