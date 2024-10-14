package main

import (
	"image/color"
	"time"

	"NBodySim/internal/nbody"
	"NBodySim/internal/simulation2d"
	"NBodySim/internal/vectormath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const FPS = 30
const MilliInFrame = 1000.0 / FPS

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("2dSim")
	myCanvas := myWindow.Canvas()

	content := container.NewWithoutLayout()
	myCanvas.SetContent(content)

	drawer := simulation2d.NewFyneDrawer(content)

	dt := 0.0001
	engine := nbody.NewIterativeNbodyEngine()
	bodies := []*nbody.Body{
		nbody.NewBody(
			vectormath.NewVector3d(400, 300, 0),
			vectormath.NewVector3d(0, 0, 0),
			1e15,
		),
		nbody.NewBody(
			vectormath.NewVector3d(500, 800, 0),
			vectormath.NewVector3d(0, 0, 0),
			2e15,
		),
		nbody.NewBody(
			vectormath.NewVector3d(800, 500, 0),
			vectormath.NewVector3d(0, 0, 0),
			1e15,
		),
		nbody.NewBody(
			vectormath.NewVector3d(300, 600, 0),
			vectormath.NewVector3d(0, 0, 0),
			1e15,
		),
	}

	body2ds := []*simulation2d.Body2d{
		simulation2d.NewBody2d(bodies[0], 10, color.NRGBA{R: 255, G: 0, B: 0, A: 255}),
		simulation2d.NewBody2d(bodies[1], 20, color.NRGBA{R: 0, G: 255, B: 0, A: 255}),
		simulation2d.NewBody2d(bodies[2], 30, color.NRGBA{R: 0, G: 0, B: 255, A: 255}),
		simulation2d.NewBody2d(bodies[3], 15, color.NRGBA{R: 255, G: 255, B: 0, A: 255}),
		// Add more bodies here...
	}

	sim := simulation2d.NewSim2d(body2ds, engine, nbody.NewEulerSolver(nil), dt)
	solver := simulation2d.NewEulerSolver(sim)
	sim.SetSolver(solver)
	sim.Draw(drawer)

	go func(sim *simulation2d.Sim2d, drawer *simulation2d.FyneDrawer) {
		timeMoment := MilliInFrame
		for true {
			start := time.Now().UnixMilli()
			sim.UpdateUntil(timeMoment / 1000)

			drawer.Clear()
			sim.Draw(drawer)
			end := time.Now().UnixMilli()
			doneIn := end - start
			m := MilliInFrame
			if doneIn < int64(m) {
				time.Sleep(time.Duration((int64(m) - doneIn) * 1e6))
			}
			timeMoment += m
			drawer.Refresh()
		}
	}(sim, drawer)

	myWindow.Resize(fyne.NewSize(1000, 1000))
	myWindow.ShowAndRun()
}
