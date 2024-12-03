package cutter

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/transform"
	"math"
	"testing"
)

var TestCube *object.PolygonObject
var TestCubeCreated bool
var TestCubeFile = "../.././models/6_hexahedron.obj"

func GetTestCube() (*object.PolygonObject, error) {
	if TestCubeCreated {
		c, _ := TestCube.Clone().(*object.PolygonObject)
		return c, nil
	}
	read, err := reader.NewObjReader(TestCubeFile)
	if err != nil {
		return nil, err
	}
	dir := builder.NewPolygonObjectDirector(&builder.InnerNormalBuilderFactory{}, read)
	cube, err := dir.Construct()
	if err != nil {
		return nil, err
	}
	TestCube, _ = cube.(*object.PolygonObject)
	TestCubeCreated = true
	c, _ := TestCube.Clone().(*object.PolygonObject)
	return c, nil
}

func TestCutCube(t *testing.T) {
	cube, err := GetTestCube()
	if err != nil {
		t.Error(err)
		return
	}
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewBackwardsCutter(camera)
	cutter.VisitPolygonObject(cube)
	if len(cube.GetPolygons()) > 10 {
		t.Errorf("Cube should have less than 10 polygons after cutting, got %v", cube.GetPolygons())
		return
	}
}

func TestCutCube2(t *testing.T) {
	cube, err := GetTestCube()
	if err != nil {
		t.Error(err)
		return
	}
	cube.Transform(transform.NewRotateAction(vector.NewVector3d(0, math.Pi/4, 0)))
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewBackwardsCutter(camera)
	cutter.VisitPolygonObject(cube)
	if len(cube.GetPolygons()) > 8 {
		t.Errorf("Cube should have less than 8 polygons after cutting, got %v", cube.GetPolygons())
		return
	}
}

func TestCutCube3(t *testing.T) {
	cube, err := GetTestCube()
	if err != nil {
		t.Error(err)
		return
	}
	cube.Transform(transform.NewRotateAction(vector.NewVector3d(math.Pi/4, math.Pi/4, 0)))
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewBackwardsCutter(camera)
	cutter.VisitPolygonObject(cube)
	if len(cube.GetPolygons()) > 6 {
		t.Errorf("Cube should have less than 6 polygons after cutting, got %v", cube.GetPolygons())
		return
	}
}

func TestCameraVisit(t *testing.T) {
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewBackwardsCutter(camera)
	camera2 := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, -1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter.VisitCamera(camera2)
	if cutter.camera != camera2 {
		t.Error("Expected camera to be updated")
		return
	}
}
