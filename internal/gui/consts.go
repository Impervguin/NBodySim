package gui

import (
	"NBodySim/internal/mathutils/vector"
	"image/color"
	"time"
)

// Window
const InitWindowWidth = 1200
const InitWindowHeight = 800
const CanvasWidth = 800
const CanvasHeight = 800

var BackgroundColor = color.RGBA{137, 142, 140, 255}

// Camera controls
const CameraRotateAngle = 10
const ZoomCameraLength = 5

// Conveyer Choose
const LightlessConvButton = "Без освещения и теней"
const LightConvButton = "С освещением, без теней"
const ShadowConvButton = "С освещением и тенями"

// Model Choose

const TetraedrModelButton = "Тетраэдр"
const CubeModelButton = "Гексаэдр"
const OctahedronModelButton = "Октаэдр"
const DodecahedronModelButton = "Додекаэдр"
const IcosahedronModelButton = "Икосаэдр"

const TetraedrModelFile = "./models/4_tetrahedron.obj"
const CubeModelFile = "./models/6_hexahedron.obj"
const OctahedronModelFile = "./models/8_octahedron.obj"
const DodecahedronModelFile = "./models/12_dodecahedron.obj"
const IcosahedronModelFile = "./models/20_icosahedron.obj"

// Mass
const MassMultiplier = 1.e9

// light
var DefaultLightColor = color.White
var DefaultLightPosition = *vector.NewVector3d(0, 10, 0)

const LightObjectScale = 0.1

// Animation
const ScreenDrawWait = time.Millisecond * 34
const SimulationTimePerFrame = time.Millisecond * 34
