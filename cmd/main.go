package main

import (
	"NBodySim/internal/nbody"
	"NBodySim/internal/vectormath"
	"fmt"
)

func main() {
	nb := nbody.NewNBody([]*nbody.Body{
		{
			Position: vectormath.NewVector3d(0, 0, 0),
			Velocity: vectormath.NewVector3d(0, 0, 0),
			Mass:     1e15,
		},
		{
			Position: vectormath.NewVector3d(100, 0, 0),
			Velocity: vectormath.NewVector3d(0, 0, 0),
			Mass:     1e15,
		},
	})
	sim := nbody.NewNBodySim(nb, nbody.NewEulerSolver(nb), nbody.NewIterativeNbodyEngine(), 0.1)
	fmt.Println(sim.GetNBody())
	for i := 0; i < 100; i++ {
		err := sim.SolveStep()
		if err != nil {
			panic(err)
		}
		for _, body := range sim.GetNBody().GetBodies() {
			fmt.Print(body.ToString(), " ")
		}
		fmt.Println()
	}

}
