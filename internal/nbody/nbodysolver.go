package nbody

import (
	"NBodySim/internal/mathutils/vector"
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

func Body2Force(target *PhysBody, influencer *PhysBody) *vector.Vector3d {
	distance := vector.SubtractVectors(&influencer.Position, &target.Position)
	distanceSquared := distance.Square()
	force := vector.NormalizeVector(distance)
	force.MultiplyScalar(G * influencer.Mass * target.Mass / distanceSquared)
	return force
}

func NewEulerSolver() *EulerSolver {
	return &EulerSolver{past: make([]PhysBody, 0)}
}

func (es *EulerSolver) CalculateBody(body *PhysBody, dt float64) (*PhysBody, error) {
	body.Position.Add(vector.MultiplyVectorScalar(&body.Velocity, dt))
	force := vector.NewVector3d(0, 0, 0)
	for _, influencer := range es.past {
		if influencer.Id != body.Id {
			force.Add(Body2Force(body, &influencer))
		}
	}
	body.Velocity.Add(vector.MultiplyVectorScalar(force, dt/body.Mass))
	return body, nil
}

func (es *EulerSolver) Reset() {
	es.past = make([]PhysBody, 0)
}

func (es *EulerSolver) UpdateSelf(cur []PhysBody) {
	es.past = make([]PhysBody, len(cur))
	copy(es.past, cur)
}

// type RungeKutta4Solver struct {
// 	past []PhysBody
// }

// func NewRungeKutta4Solver() *RungeKutta4Solver {
// 	return &RungeKutta4Solver{past: make([]PhysBody, 0)}
// }

// func (rk4 *RungeKutta4Solver) CalculateBody(body *PhysBody, dt float64) (*PhysBody, error) {
// 	vel1 := body.Velocity.Copy()
// 	force1 := vector.NewVector3d(0, 0, 0)
// 	for _, influencer := range rk4.past {
// 		if influencer.Id != body.Id {
// 			force1.Add(Body2Force(body, &influencer))
// 		}
// 	}

// }

// func (rk4 *RungeKutta4Solver) Reset() {
// 	rk4.past = make([]PhysBody, 0)
// }

// func (rk4 *RungeKutta4Solver) UpdateSelf(cur []PhysBody) {
// 	rk4.past = make([]PhysBody, len(cur))
// 	copy(rk4.past, cur)
// }
