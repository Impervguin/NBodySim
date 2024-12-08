package main

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/nbody"
	"fmt"
	"math"
	"strings"
	"sync"
)

type ResearchUnit struct {
	TimeMax    float64
	Dt         float64
	MinPercent float64
	MaxPercent float64
	SumPercent float64
	Count      int
	Nb         *nbody.NBody
	Config     int
}

func NewResearchUnit(tm, dt float64, nb *nbody.NBody, config int) *ResearchUnit {
	return &ResearchUnit{
		TimeMax:    tm,
		Dt:         dt,
		MinPercent: math.Inf(1),
		MaxPercent: math.Inf(-1),
		SumPercent: 0.,
		Count:      0,
		Nb:         nb,
		Config:     config,
	}
}

func (ru *ResearchUnit) DoResearch() {
	envis := nbody.NewEnergyVisitor()
	ru.Nb.Accept(envis)
	initEnergy := envis.GetTotalEnergy()
	for t := ru.Dt; t < ru.TimeMax; t += ru.Dt {
		ru.Nb.SolveStep(ru.Dt)
		envis = nbody.NewEnergyVisitor()
		envis.VisitNBody(ru.Nb)
		energy := envis.GetTotalEnergy()
		diff := energy - initEnergy
		diffPercent := math.Abs(diff / initEnergy * 100)
		ru.SumPercent += diffPercent
		ru.Count++
		ru.MinPercent = math.Min(ru.MinPercent, diffPercent)
		ru.MaxPercent = math.Max(ru.MaxPercent, diffPercent)
	}
}

func (ru *ResearchUnit) LatexTableString() string {
	return fmt.Sprintf("%.f & %.4g & %d & %.4g & %.4g & %.4g \\\\ \\hline", ru.TimeMax, ru.Dt, ru.Config, ru.MinPercent, ru.MaxPercent, ru.SumPercent/float64(ru.Count))
}

func (ru *ResearchUnit) LatexConfigString() string {
	bodies, _ := ru.Nb.GetBodies()
	positions := make([]string, 1, len(bodies)+1)
	positions[0] = "\\textbf{Радиус-вектор}"

	for _, b := range bodies {
		positions = append(positions, fmt.Sprintf("$\\begin{pmatrix} %.f \\\\ %.f \\\\ %.f \\end{pmatrix}$", b.GetPosition().X, b.GetPosition().Y, b.GetPosition().Z))
	}
	pos := strings.Join(positions, " & ")
	pos += "\\\\ \\hline"
	return pos
}

func ResearchCycle1() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 30), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(30, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -30), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-30, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	return NewResearchUnit(300, 0.01, nb, 1)
}

func ResearchCycle2() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 30), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(30, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -30), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-30, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	return NewResearchUnit(300, 0.001, nb, 1)
}

func ResearchCycle3() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 30), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(30, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -30), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-30, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	return NewResearchUnit(300, 0.0001, nb, 1)
}

func ResearchCollision1() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(-5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	nb.AddBody(body1)
	nb.AddBody(body2)

	return NewResearchUnit(11, 0.001, nb, 2)
}

func ResearchCollision2() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(-5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	nb.AddBody(body1)
	nb.AddBody(body2)

	return NewResearchUnit(11, 0.0001, nb, 2)
}

func ResearchCollision3() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(-5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	nb.AddBody(body1)
	nb.AddBody(body2)

	return NewResearchUnit(11, 0.00001, nb, 2)
}

func ResearchCollision4() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(-5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	nb.AddBody(body1)
	nb.AddBody(body2)

	return NewResearchUnit(11, 0.000001, nb, 2)
}

func ResearchCollision5() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(-5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(5, 0, 0), *vector.NewVector3d(0, 0, 0), 10e10)
	nb.AddBody(body1)
	nb.AddBody(body2)

	return NewResearchUnit(11, 0.0000001, nb, 2)
}

func ResearchSphere1() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 0), 1e12)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(25, 0, 0), *vector.NewVector3d(0, 0, -1.6334), 1)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(-25, 0, 0), *vector.NewVector3d(0, 0, 1.6334), 1)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(0, 25, 0), *vector.NewVector3d(0, 0, 1.6334), 1)
	body5 := nbody.NewOnlyBody(*vector.NewVector3d(30, 0, 0), *vector.NewVector3d(0, 1.491, 0), 1)
	body6 := nbody.NewOnlyBody(*vector.NewVector3d(45, 0, 0), *vector.NewVector3d(0, 0, 1.2174), 1)
	body7 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -30), *vector.NewVector3d(0, 1.491, 0), 1)
	body8 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 40), *vector.NewVector3d(0, 1.2913, 0), 1)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	nb.AddBody(body5)
	nb.AddBody(body6)
	nb.AddBody(body7)
	nb.AddBody(body8)

	return NewResearchUnit(10000, 0.01, nb, 3)
}

func ResearchSphere2() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 0), 1e12)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(25, 0, 0), *vector.NewVector3d(0, 0, -1.6334), 1)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(-25, 0, 0), *vector.NewVector3d(0, 0, 1.6334), 1)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(0, 25, 0), *vector.NewVector3d(0, 0, 1.6334), 1)
	body5 := nbody.NewOnlyBody(*vector.NewVector3d(30, 0, 0), *vector.NewVector3d(0, 1.491, 0), 1)
	body6 := nbody.NewOnlyBody(*vector.NewVector3d(45, 0, 0), *vector.NewVector3d(0, 0, 1.2174), 1)
	body7 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -30), *vector.NewVector3d(0, 1.491, 0), 1)
	body8 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 40), *vector.NewVector3d(0, 1.2913, 0), 1)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	nb.AddBody(body5)
	nb.AddBody(body6)
	nb.AddBody(body7)
	nb.AddBody(body8)

	return NewResearchUnit(10000, 0.001, nb, 3)
}

func ResearchSphere3() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 0), 1e12)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(25, 0, 0), *vector.NewVector3d(0, 0, -1.6334), 1)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(-25, 0, 0), *vector.NewVector3d(0, 0, 1.6334), 1)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(0, 25, 0), *vector.NewVector3d(0, 0, 1.6334), 1)
	body5 := nbody.NewOnlyBody(*vector.NewVector3d(30, 0, 0), *vector.NewVector3d(0, 1.491, 0), 1)
	body6 := nbody.NewOnlyBody(*vector.NewVector3d(45, 0, 0), *vector.NewVector3d(0, 0, 1.2174), 1)
	body7 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -30), *vector.NewVector3d(0, 1.491, 0), 1)
	body8 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 40), *vector.NewVector3d(0, 1.2913, 0), 1)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	nb.AddBody(body5)
	nb.AddBody(body6)
	nb.AddBody(body7)
	nb.AddBody(body8)

	return NewResearchUnit(10000, 0.0001, nb, 3)
}

func main() {
	rus := []*ResearchUnit{
		ResearchCycle1(),
		ResearchCycle2(),
		ResearchCycle3(),
		ResearchCollision1(),
		ResearchCollision2(),
		ResearchCollision3(),
		ResearchCollision4(),
		ResearchCollision5(),
		ResearchSphere1(),
		ResearchSphere2(),
		ResearchSphere3(),
	}

	wg := sync.WaitGroup{}
	for _, ru := range rus {
		wg.Add(1)
		go func() { ru.DoResearch(); wg.Done() }()
	}
	wg.Wait()
	for _, ru := range rus {
		fmt.Printf("%s\n", ru.LatexTableString())
	}
	// for _, ru := range rus {
	// 	fmt.Printf("%s\n\n", ru.LatexConfigString())
	// }
}
