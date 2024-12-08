package nbody

import (
	"NBodySim/internal/mathutils/vector"
)

type PotentialVisitor struct {
	potentialEnergy float64
}

func NewPotentialVisitor() *PotentialVisitor {
	return &PotentialVisitor{
		potentialEnergy: 0,
	}
}

func (v *PotentialVisitor) VisitBody(body Body) {
}

func (v *PotentialVisitor) VisitNBody(nbody *NBody) {
	for id1, b1 := range nbody.nbody {
		for id2, b2 := range nbody.nbody {
			if id1 < id2 {
				p1 := b1.GetPosition()
				p2 := b2.GetPosition()
				distance := vector.Length(vector.SubtractVectors(&p1, &p2))
				// fmt.Println(-b1.GetMass() * b2.GetMass() / distance)
				v.potentialEnergy += -G * b1.GetMass() * b2.GetMass() / distance
			}
		}
	}
}

func (v *PotentialVisitor) GetPotentialEnergy() float64 {
	return v.potentialEnergy
}
