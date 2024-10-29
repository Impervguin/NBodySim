package nbody

import (
	"NBodySim/internal/vectormath"
)

const (
	G float64 = 6.67430e-11 // Newton's gravitational constant
)

type NBodySolver interface {
	CalculateBody(body *PhysBody, dt float64) (*PhysBody, error)
	UpdateSelf(cur []PhysBody)
	Reset()
}

type EulerSolver struct {
	past []PhysBody
}

func Body2Force(target *PhysBody, influencer *PhysBody) *vectormath.Vector3d {
	distance := vectormath.SubtractVectors(&influencer.Position, &target.Position)
	distanceSquared := distance.Square()
	force := vectormath.NormalizeVector(distance)
	force.MultiplyScalar(G * influencer.Mass * target.Mass / distanceSquared)
	return force
}

func NewEulerSolver() *EulerSolver {
	return &EulerSolver{past: make([]PhysBody, 0)}
}

func (es *EulerSolver) CalculateBody(body *PhysBody, dt float64) (*PhysBody, error) {
	body.Position.Add(vectormath.MultiplyVectorScalar(&body.Velocity, dt))
	force := vectormath.NewVector3d(0, 0, 0)
	for _, influencer := range es.past {
		if influencer.Id != body.Id {
			force.Add(Body2Force(body, &influencer))
		}
	}
	body.Velocity.Add(vectormath.MultiplyVectorScalar(force, dt/body.Mass))
	return body, nil
}

func (es *EulerSolver) Reset() {
	es.past = make([]PhysBody, 0)
}

func (es *EulerSolver) UpdateSelf(cur []PhysBody) {
	es.past = make([]PhysBody, len(cur))
	copy(es.past, cur)
}
