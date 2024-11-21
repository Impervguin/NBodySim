package gui

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/conveyer"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/reader"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"NBodySim/internal/zmapper/mapper"
	"NBodySim/internal/zmapper/objectdrawer"
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NBodyApp struct {
	napp          fyne.App
	win           fyne.Window
	canvas        *canvas.Raster
	canvasImage   *image.Image
	cameraMan     CameraManager
	sim           simulation.Simulation
	conv          conveyer.RenderConveyer
	lightConv     conveyer.RefactoredSimulationConveyer
	lightlessConv conveyer.SimplestSimulationConveyer
	width         float64
	height        float64
}

const InitWindowWidth = 1200
const InitWindowHeight = 800
const CanvasWidth = 800
const CanvasHeight = 800

const CameraRotateAngle = 10
const ZoomCameraLength = 5

func NewNBodyApp() *NBodyApp {
	app := app.New()
	win := app.NewWindow("N-Body Simulation")

	width := 1200.
	height := 800.

	win.Resize(fyne.NewSize(float32(width), float32(height)))
	return &NBodyApp{
		napp:   app,
		win:    win,
		width:  width,
		height: height,
	}
}

func (na *NBodyApp) initSimulation() {
	read, _ := reader.NewObjReader("./models/6_hexahedron.obj")
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube, err := dir.Construct()
	if err != nil {
		panic(err)
	}

	// read, _ = reader.NewObjReader("./models/6_hexahedron.obj")
	// dir = builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube2, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	cube2.Transform(transform.NewMoveAction(vector.NewVector3d(5, 0, 0)))
	na.sim.AddObject(cube, *vector.NewVector3d(0, 0, 0), 1)
	na.sim.AddObject(cube2, *vector.NewVector3d(0, 0, 0), 1)
	na.cameraMan.MoveCamera(0, 0, -15)
	na.conv.Convey()
	// na.cameraMan.ZoomCamera(-15)
}

func (na *NBodyApp) updateCanvas() {
	na.conv.Convey()
	na.canvas.Refresh()
}

func (na *NBodyApp) createLayout() {
	time.Sleep(time.Second)
	na.width, na.height = float64(na.width)*float64(na.win.Canvas().Scale()), float64(na.height)*float64(na.win.Canvas().Scale())

	cw, ch := CanvasWidth*na.win.Canvas().Scale(), CanvasHeight*na.win.Canvas().Scale()

	na.sim = *simulation.NewSimulation()
	// na.cameraMan = NewFreeCameraManager(na.sim.GetCamera())
	na.cameraMan = NewCentricCameraManager(na.sim.GetCamera())

	na.lightConv = *conveyer.NewRefactoredSimulationConveyer(
		objectdrawer.NewParallelWithoutLightsDrawerFabric(
			mapper.NewParallelZmapperWithNormalsFabric(int(cw), int(ch), color.Black, &buffers.DepthBufferNullFabric{}),
			approximator.NewFlatNormalApproximatorFabric()),
		&na.sim,
	)
	na.lightlessConv = *conveyer.NewSimplestSimulationConveyer(
		objectdrawer.NewParallelWithoutLightsDrawerFabric(
			mapper.NewParallelZmapperWithNormalsFabric(int(cw), int(ch), color.Black, &buffers.DepthBufferNullFabric{}),
			approximator.NewFlatNormalApproximatorFabric()),
		&na.sim,
	)

	na.conv = &na.lightlessConv
	na.initSimulation()

	// TODO: остальные конвейеры

	/*
		Канвас
	*/
	im := na.conv.GetImage()
	na.canvasImage = &im

	na.canvas = canvas.NewRasterFromImage(*na.canvasImage)
	na.canvas.SetMinSize(fyne.NewSize(CanvasWidth, CanvasHeight))

	/*
	   Таб сцены
	*/

	startButton := widget.NewButton("Запустить", func() {})
	stopButton := widget.NewButton("Остановить", func() {})
	simControls := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer())

	// Камеры работают неправильно
	leftButton := widget.NewButton("<-", func() { na.cameraMan.RotateRight(-CameraRotateAngle); na.updateCanvas() })
	rightButton := widget.NewButton("->", func() { na.cameraMan.RotateRight(CameraRotateAngle); na.updateCanvas() })
	upButton := widget.NewButton("^", func() { na.cameraMan.RotateUp(CameraRotateAngle); na.updateCanvas() })
	downButton := widget.NewButton("v", func() { na.cameraMan.RotateUp(-CameraRotateAngle); na.updateCanvas() })

	zoomFarButton := widget.NewButton("Дальше", func() { na.cameraMan.MoveCamera(0, 0, -ZoomCameraLength); na.updateCanvas() })
	zoomNearButton := widget.NewButton("Ближе", func() { na.cameraMan.MoveCamera(0, 0, ZoomCameraLength); na.updateCanvas() })
	cameraControls := container.NewVBox(
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Управление камерой"), layout.NewSpacer()),
		container.NewHBox(
			layout.NewSpacer(),
			container.NewVBox(
				container.NewHBox(layout.NewSpacer(), widget.NewLabel("Поворот"), layout.NewSpacer()),
				container.NewBorder(upButton, downButton, leftButton, rightButton),
			),
			layout.NewSpacer(),
			container.NewVBox(
				container.NewHBox(layout.NewSpacer(), widget.NewLabel("Приближение"), layout.NewSpacer()),
				container.NewBorder(zoomNearButton, zoomFarButton, nil, nil),
			),
			layout.NewSpacer(),
		),
	)

	sceneTabContainer := container.NewVBox(layout.NewSpacer(), simControls, layout.NewSpacer(), cameraControls, layout.NewSpacer())

	tabs := container.NewAppTabs(
		container.NewTabItem("Сцена", sceneTabContainer),
		container.NewTabItem("Объекты", widget.NewLabel("Объекты")),
	)
	outerlayout := container.NewBorder(nil, nil, na.canvas, nil, tabs)
	na.win.SetContent(outerlayout)
}

func (na *NBodyApp) StartApp() {
	go na.createLayout()
	na.win.ShowAndRun()
}
