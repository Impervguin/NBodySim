package gui

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/conveyer"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"NBodySim/internal/zmapper/mapper"
	"NBodySim/internal/zmapper/objectdrawer"
	"context"
	"fmt"
	"image/color"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NBodyApp struct {
	napp          fyne.App
	win           fyne.Window
	canvas        *canvas.Raster
	canvasBox     *fyne.Container
	cameraMan     CameraManager
	sim           simulation.Simulation
	conv          conveyer.RenderConveyer
	lightConv     conveyer.RefactoredSimulationConveyer
	lightlessConv conveyer.SimplestSimulationConveyer
	shadowConv    conveyer.RefactoredShadowSimulationConveyer
	light         *object.PointLightShadow
	lightObject   *object.PolygonObject
	width         float64
	height        float64

	chosenModelName string
	chosenModelFile string
	modelColor      color.Color
	modelColorRect  *canvas.Rectangle
	modelSize       float64
	modelMass       float64
	modelPosition   vector.Vector3d
	modelVelocity   vector.Vector3d

	createdObjects map[int64]string
	objectsSelect  *widget.Select
	selectedObject string

	lightPosition      vector.Vector3d
	lightIntensity     color.Color
	lightIntensityRect *canvas.Rectangle

	simulationContext context.Context
	simulationCancel  context.CancelFunc

	tabs      *container.AppTabs
	objectTab *container.TabItem
	lightTab  *container.TabItem

	startButton *widget.Button
	stopButton  *widget.Button
}

func NewNBodyApp() *NBodyApp {
	app := app.New()
	win := app.NewWindow("N-Body Simulation")

	width := 1200.
	height := 800.

	win.Resize(fyne.NewSize(float32(width), float32(height)))
	return &NBodyApp{
		napp:           app,
		win:            win,
		width:          width,
		height:         height,
		createdObjects: make(map[int64]string),
	}
}

func (na *NBodyApp) initSimulation() {
	na.createObject()
	na.cameraMan.MoveCamera(0, 0, -15)

	na.light = object.NewPointLightShadow(color.White, *vector.NewVector3d(0, 0, 0))
	na.sim.AddLight(na.light)

	// na.lightObject = object.NewPolygonObject()
	read, err := reader.NewObjReader(TetraedrModelFile)
	if err != nil {
		panic(err)
	}
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	obj, err := dir.Construct()
	if err != nil {
		panic(err)
	}

	obj.Transform(transform.NewScaleAction(vector.NewVector3d(LightObjectScale, LightObjectScale, LightObjectScale)))
	na.lightObject = obj.(*object.PolygonObject)
	na.sim.AddImaginaryObject(obj)

	na.updateLight()
}

func (na *NBodyApp) updateCanvas() {
	na.conv.Convey()
	na.canvas.Refresh()
}

func (na *NBodyApp) changeConveyer(value string) {
	switch value {
	case LightlessConvButton:
		na.conv = &na.lightlessConv
	case LightConvButton:
		na.conv = &na.lightConv
	default:
		na.conv = &na.shadowConv
	}
	na.canvas = canvas.NewRasterFromImage(na.conv.GetImage())
	na.canvas.SetMinSize(fyne.NewSize(CanvasWidth, CanvasHeight))
	na.canvasBox.RemoveAll()
	na.canvasBox.Add(na.canvas)

	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) chooseModel(model string) {
	switch model {
	case TetraedrModelButton:
		na.chosenModelFile = TetraedrModelFile
	case CubeModelButton:
		na.chosenModelFile = CubeModelFile
	case OctahedronModelButton:
		na.chosenModelFile = OctahedronModelFile
	case DodecahedronModelButton:
		na.chosenModelFile = DodecahedronModelFile
	case IcosahedronModelButton:
		na.chosenModelFile = IcosahedronModelFile
	}
	na.chosenModelName = model
}

func (na *NBodyApp) setModelColor(c color.Color) {
	na.modelColor = c
	na.modelColorRect.FillColor = na.modelColor
	na.modelColorRect.Refresh()
}

func (na *NBodyApp) setLightColor(c color.Color) {
	na.lightIntensity = c
	na.lightIntensityRect.FillColor = na.lightIntensity
	na.lightIntensityRect.Refresh()
}

func (na *NBodyApp) createObject() {
	read, err := reader.NewObjReader(na.chosenModelFile)
	if err != nil {
		dialog.NewError(err, na.win).Show()
		return
	}
	dir := builder.NewPolygonObjectDirector(&builder.InnerNormalBuilderFactory{}, read)
	obj, err := dir.Construct()
	if err != nil {
		dialog.NewError(err, na.win).Show()
		return
	}
	pobj, _ := obj.(*object.PolygonObject)
	pobj.SetColor(na.modelColor)
	obj.Transform(transform.NewScaleAction(vector.NewVector3d(na.modelSize, na.modelSize, na.modelSize)))
	obj.Transform(transform.NewMoveAction(&na.modelPosition))
	na.sim.AddObject(obj, na.modelVelocity, na.modelMass*MassMultiplier)

	na.createdObjects[obj.GetId()] = na.chosenModelName
	na.updateObjectSelect()

	na.updateCanvas()
}

func (na *NBodyApp) deleteObject() {
	objString := na.selectedObject
	if objString == "" {
		dialog.NewError(fmt.Errorf("Не выбран объект"), na.win).Show()
		return
	}
	objId, err := strconv.ParseInt(strings.Split(objString, ":")[0], 10, 64)
	if err != nil {
		dialog.NewError(err, na.win).Show()
		return
	}
	err = na.sim.RemoveObject(objId)
	if err != nil {
		dialog.NewError(err, na.win).Show()
		return
	}
	delete(na.createdObjects, objId)
	na.updateObjectSelect()
	na.objectsSelect.ClearSelected()
	na.updateCanvas()
}

func (na *NBodyApp) updateObjectSelect() {
	options := make([]string, 0, len(na.createdObjects))
	for id, name := range na.createdObjects {
		options = append(options, fmt.Sprintf("%d: %s", id, name))
	}
	na.objectsSelect.SetOptions(options)
}

func (na *NBodyApp) updateLight() {
	na.lightObject.SetColor(na.lightIntensity)
	na.lightObject.MoveCenterTo(na.lightPosition)
	na.light.SetIntensity(na.lightIntensity)
	na.light.SetPosition(na.lightPosition)

	na.updateCanvas()
}

func (na *NBodyApp) cameraRotateUp() {
	na.cameraMan.RotateUp(CameraRotateAngle)
	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) cameraRotateDown() {
	na.cameraMan.RotateUp(-CameraRotateAngle)
	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) cameraRotateLeft() {
	na.cameraMan.RotateRight(-CameraRotateAngle)
	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) cameraRotateRight() {
	na.cameraMan.RotateRight(CameraRotateAngle)
	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) cameraMoveForward() {
	na.cameraMan.MoveCamera(0, 0, ZoomCameraLength)
	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) cameraMoveBack() {
	na.cameraMan.MoveCamera(0, 0, -ZoomCameraLength)
	if na.simulationContext.Err() != nil {
		na.updateCanvas()
	}
}

func (na *NBodyApp) createLayout() {
	time.Sleep(time.Second)
	na.width, na.height = float64(na.width)*float64(na.win.Canvas().Scale()), float64(na.height)*float64(na.win.Canvas().Scale())

	cw, ch := CanvasWidth*na.win.Canvas().Scale(), CanvasHeight*na.win.Canvas().Scale()

	na.sim = *simulation.NewSimulation()
	na.sim.SetCamera(
		object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, -1, 0), 1, 1, 1),
	)
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
	na.shadowConv = *conveyer.NewRefactoredShadowSimulationConveyer(
		objectdrawer.NewParallelWithoutLightsDrawerFabric(
			mapper.NewParallelZmapperWithNormalsFabric(int(cw), int(ch), color.Black, &buffers.DepthBufferNullFabric{}),
			approximator.NewFlatNormalApproximatorFabric(),
		),
		&na.sim,
	)

	/*
		Канвас
	*/
	na.canvasBox = container.NewHBox(na.canvas)

	/*
	   Таб сцены
	*/

	startButton := widget.NewButton("Запустить", func() { na.StartSimulation() })
	na.startButton = startButton
	stopButton := widget.NewButton("Остановить", func() { na.StopSimulation() })
	na.stopButton = stopButton
	na.stopButton.Disable()
	simControls := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer(), stopButton, layout.NewSpacer())

	leftButton := widget.NewButton("<-", na.cameraRotateLeft)
	rightButton := widget.NewButton("->", na.cameraRotateRight)
	upButton := widget.NewButton("^", na.cameraRotateUp)
	downButton := widget.NewButton("v", na.cameraRotateDown)

	zoomFarButton := widget.NewButton("Дальше", na.cameraMoveBack)
	zoomNearButton := widget.NewButton("Ближе", na.cameraMoveForward)
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

	convRadio := widget.NewRadioGroup([]string{LightlessConvButton, LightConvButton, ShadowConvButton}, na.changeConveyer)
	convRadio.SetSelected(LightlessConvButton)

	convRadioBox := container.NewHBox(layout.NewSpacer(), convRadio, layout.NewSpacer())

	sceneTabContainer := container.NewVBox(layout.NewSpacer(), simControls, layout.NewSpacer(), cameraControls, layout.NewSpacer(), convRadioBox, layout.NewSpacer())

	/*
	 Таб объектов
	*/

	modelRadio := widget.NewRadioGroup([]string{TetraedrModelButton, CubeModelButton, OctahedronModelButton, DodecahedronModelButton, IcosahedronModelButton}, na.chooseModel)
	modelRadio.SetSelected(CubeModelButton)
	modelRadio.Required = true

	rect := canvas.NewRectangle(color.White)
	rect.SetMinSize(fyne.NewSize(40, 20))
	na.modelColorRect = rect
	na.setModelColor(color.White)

	modelRadioBox := container.NewHBox(layout.NewSpacer(), modelRadio, layout.NewSpacer())

	sizeInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelSize)))
	sizeInput.SetText("1")
	sizeInputBox := container.NewBorder(nil, nil,
		widget.NewLabel("Размер объекта:"),
		nil,
		sizeInput,
	)

	colorPickButton := container.NewStack(widget.NewButton("", func() {
		dialog.NewColorPicker("Цвет объекта", "Выберете цвет объекта", na.setModelColor, na.win).Show()
	}), rect)
	colorPick := container.NewHBox(widget.NewLabel("Цвет объекта:"), colorPickButton, layout.NewSpacer())

	massInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelMass)))
	massInput.SetText("1")
	massInputBox := container.NewBorder(nil, nil,
		widget.NewLabel("Масса объекта:"),
		widget.NewLabel("*10^9"),
		massInput,
	)

	xPosInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelPosition.X)))
	yPosInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelPosition.Y)))
	zPosInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelPosition.Z)))
	positionInputBox := container.NewGridWithColumns(3,
		container.NewBorder(nil, nil, widget.NewLabel("X:"), nil, xPosInput),
		container.NewBorder(nil, nil, widget.NewLabel("Y:"), nil, yPosInput),
		container.NewBorder(nil, nil, widget.NewLabel("Z:"), nil, zPosInput),
	)

	xVelInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelVelocity.X)))
	yVelInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelVelocity.Y)))
	zVelInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.modelVelocity.Z)))

	velocityInputBox := container.NewGridWithColumns(3,
		container.NewBorder(nil, nil, widget.NewLabel("Vx:"), nil, xVelInput),
		container.NewBorder(nil, nil, widget.NewLabel("Vy:"), nil, yVelInput),
		container.NewBorder(nil, nil, widget.NewLabel("Vz:"), nil, zVelInput),
	)

	createObjectButton := widget.NewButton("Создать объект", na.createObject)
	createObjectButtonBox := container.NewHBox(layout.NewSpacer(), createObjectButton, layout.NewSpacer())

	na.objectsSelect = widget.NewSelect([]string{}, func(s string) {
		na.selectedObject = s
	})

	deleteObjectButton := widget.NewButton("Удалить выбранный объект", na.deleteObject)
	deleteObjectButtonBox := container.NewHBox(layout.NewSpacer(), deleteObjectButton, layout.NewSpacer())

	objectTabContainer := container.NewVBox(
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Визуальные характеристики"), layout.NewSpacer()),
		modelRadioBox,
		colorPick,
		sizeInputBox,
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Физические характеристики"), layout.NewSpacer()),
		massInputBox,
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Позиция объекта"), layout.NewSpacer()),
		positionInputBox,
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Скорость объекта"), layout.NewSpacer()),
		velocityInputBox,
		createObjectButtonBox,
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Удаление объектов"), layout.NewSpacer()),
		na.objectsSelect,
		deleteObjectButtonBox,
	)

	/*
	   Таб управления светом
	*/

	xLightPosInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.lightPosition.X)))
	yLightPosInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.lightPosition.Y)))
	zLightPosInput := widget.NewEntryWithData(binding.FloatToString(binding.BindFloat(&na.lightPosition.Z)))
	xLightPosInput.SetText(strconv.FormatFloat(DefaultLightPosition.X, 'f', 2, 64))
	yLightPosInput.SetText(strconv.FormatFloat(DefaultLightPosition.Y, 'f', 2, 64))
	zLightPosInput.SetText(strconv.FormatFloat(DefaultLightPosition.Z, 'f', 2, 64))
	lightPositionInputBox := container.NewGridWithColumns(3,
		container.NewBorder(nil, nil, widget.NewLabel("X:"), nil, xLightPosInput),
		container.NewBorder(nil, nil, widget.NewLabel("Y:"), nil, yLightPosInput),
		container.NewBorder(nil, nil, widget.NewLabel("Z:"), nil, zLightPosInput),
	)

	na.lightIntensityRect = canvas.NewRectangle(DefaultLightColor)
	na.lightIntensity = DefaultLightColor

	lightColorPickButton := container.NewStack(widget.NewButton("", func() {
		dialog.NewColorPicker("Цвет освещения", "Выберете цвет света", na.setLightColor, na.win).Show()
	}), na.lightIntensityRect)
	lightColorPick := container.NewHBox(widget.NewLabel("Цвет освещения:"), lightColorPickButton, layout.NewSpacer())

	lightUpdateButton := widget.NewButton("Обновить освещение", func() { na.updateLight() })
	lightUpdateButtonBox := container.NewHBox(layout.NewSpacer(), lightUpdateButton, layout.NewSpacer())

	lightTabContainer := container.NewVBox(
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Интенсивность источника света"), layout.NewSpacer()),
		lightColorPick,
		container.NewHBox(layout.NewSpacer(), widget.NewLabel("Позиция источника света"), layout.NewSpacer()),
		lightPositionInputBox,
		lightUpdateButtonBox,
	)

	na.objectTab = container.NewTabItem("Объекты", objectTabContainer)
	na.lightTab = container.NewTabItem("Освещение", lightTabContainer)

	tabs := container.NewAppTabs(
		container.NewTabItem("Сцена", sceneTabContainer),
		na.objectTab,
		na.lightTab,
	)
	na.tabs = tabs
	outerlayout := container.NewBorder(nil, nil, na.canvasBox, nil, tabs)
	na.win.SetContent(outerlayout)
	na.initSimulation()
}

func (na *NBodyApp) StartSimulation() {
	na.simulationContext, na.simulationCancel = context.WithCancel(context.Background())
	na.tabs.DisableItem(na.objectTab)
	na.tabs.DisableItem(na.lightTab)
	na.startButton.Disable()
	na.stopButton.Enable()
	go func() {
		fmt.Println("Started")
		for {
			select {
			case <-na.simulationContext.Done():
				fmt.Println("Stopped")
				return
			default:
				na.sim.UpdateFor(float64(SimulationTimePerFrame) / float64(time.Second))
				na.updateCanvas()
				time.Sleep(ScreenDrawWait)
			}
		}
	}()
}

func (na *NBodyApp) StopSimulation() {
	na.simulationCancel()
	na.tabs.EnableItem(na.objectTab)
	na.tabs.EnableItem(na.lightTab)
	na.startButton.Enable()
	na.stopButton.Disable()
}

func (na *NBodyApp) StartApp() {
	na.simulationContext, na.simulationCancel = context.WithCancel(context.Background())
	na.simulationCancel()
	go na.createLayout()
	na.win.ShowAndRun()
}
