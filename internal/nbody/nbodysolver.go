package nbody

import "NBodySim/internal/vectormath"

const (
	G float64 = 6.67430e-11 // Newton's gravitational constant
)

type NBodySolver interface {
	CalculateBody(body *Body, dt float64) (*Body, error)
	UpdateSelf(nbody *NBody)
}

type EulerSolver struct {
	nbody *NBody
}

func Body2Force(target *Body, influencer *Body) *vectormath.Vector3d {
	distance := vectormath.SubtractVectors(influencer.Position, target.Position)
	distanceSquared := distance.Square()
	force := vectormath.NormalizeVector(distance)
	force.MultiplyScalar(G * influencer.Mass * target.Mass / distanceSquared)
	return force
}

func NewEulerSolver(nbody *NBody) *EulerSolver {
	return &EulerSolver{nbody: nbody}
}

func (es *EulerSolver) CalculateBody(body *Body, dt float64) (*Body, error) {
	ubody := body.Copy()
	ubody.Position.Add(vectormath.MultiplyVectorScalar(body.Velocity, dt))
	force := vectormath.NewVector3d(0, 0, 0)
	for _, influencer := range es.nbody.bodies {
		if influencer != body {
			force.Add(Body2Force(body, influencer))
		}
	}
	ubody.Velocity.Add(vectormath.MultiplyVectorScalar(force, dt/body.Mass))
	return ubody, nil
}

func (es *EulerSolver) UpdateSelf(nbody *NBody) {
	es.nbody = nbody
}
