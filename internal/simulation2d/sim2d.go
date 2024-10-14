package simulation2d

import (
	"NBodySim/internal/nbody"
)

type Sim2d struct {
	nbsim      *nbody.NBodySim
	bodies     map[int64]*Body2d
	dt         float64
	timeMoment float64
}

func NewSim2d(bodies []*Body2d, engine nbody.NBodyEngine, solver nbody.NBodySolver, dt float64) *Sim2d {
	nbodies := make([]*nbody.Body, len(bodies), len(bodies))

	for i, body := range bodies {
		nbodies[i] = body.body
	}
	nbodysim := nbody.NewNBodySim(nbody.NewNBody(nbodies), solver, engine, dt)

	bodies2d := make(map[int64]*Body2d, len(bodies))
	for _, body := range bodies {
		bodies2d[body.body.Id] = body
	}
	return &Sim2d{nbsim: nbodysim,
		bodies:     bodies2d,
		dt:         dt,
		timeMoment: 0.0,
	}
}

func NewEulerSolver(sim *Sim2d) *nbody.EulerSolver {
	return nbody.NewEulerSolver(sim.nbsim.GetNBody())
}

func (s *Sim2d) SetSolver(solver nbody.NBodySolver) {
	s.nbsim.SetSolver(solver)
}

func (s *Sim2d) Update() {
	s.nbsim.SolveStep()
	s.timeMoment += s.dt
}

func (s *Sim2d) UpdateUntil(timeMoment float64) {
	for s.timeMoment < timeMoment {
		s.Update()
	}
}

func (s *Sim2d) Draw(drawer Sim2dDrawer) {
	drawer.Clear()
	for _, body := range s.bodies {
		// fmt.Println(body.body.Position)
		xc := int64(body.body.Position.X)
		yc := int64(body.body.Position.Y)
		drawer.DrawCircle(body.radius, body.color, xc, yc)
	}
	drawer.Refresh()
}
